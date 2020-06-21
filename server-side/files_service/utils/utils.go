package utils

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	datastructures "../data_structures"
	"github.com/go-redis/redis"
)

var redisClient *redis.Client
var ctx context.Context

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
	})

	ctx = redisClient.Context()
}

// GetFilesFromDisk returns the files for a workspace from the disk
func GetFilesFromDisk(ws datastructures.Workspace) ([]datastructures.File, error) {
	path := ws.ToString()
	sep := string(os.PathSeparator)
	pathLen := len(strings.Split(path, sep))

	var files []datastructures.File

	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			var data []byte

			if !info.IsDir() {
				if data, err = ioutil.ReadFile(path); err != nil {
					return err
				}
			}

			pathList := strings.Split(path, sep)[pathLen:]

			f := datastructures.File{
				Path:  pathList,
				Data:  string(data),
				IsDir: info.IsDir(),
			}

			if len(pathList) > 0 {
				files = append(files, f)
			}

			return nil
		})

	return files, err
}

// CommitFilesToDisk copies to disk all the files for a workspace
// If a file is marked as deleted , then delete it
func CommitFilesToDisk(ws datastructures.Workspace) error {
	res, err := redisClient.HGetAll(ctx, ws.ToString()).Result()

	if err != nil {
		return err
	}

	for name, entry := range res {
		var cacheEntry datastructures.CacheEntry

		if err = json.Unmarshal([]byte(entry), &cacheEntry); err != nil {
			return err
		}

		path := filepath.Join(ws.ToString(), name)

		if cacheEntry.Deleted {
			if err := os.RemoveAll(path); err != nil {
				return err
			}
		} else {
			if !cacheEntry.IsDir {
				if err := ioutil.WriteFile(path, []byte(cacheEntry.Data), 0666); err != nil {
					return err
				}
			} else {
				if err := os.MkdirAll(path, 0666); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// IsWorkspaceInCache verifies if files of an workspace are in cache
func IsWorkspaceInCache(ws datastructures.Workspace) (bool, error) {
	exists, err := redisClient.Exists(ctx, ws.ToString()).Result()

	if err != nil {
		return false, err
	}

	return !(exists == 0), nil
}

// IsFileInCache verifies if a specific file of an workspace is in cache
func IsFileInCache(ws datastructures.Workspace, path []string) (bool, error) {
	var exists bool
	var cacheEntry datastructures.CacheEntry

	if exists, err := redisClient.HExists(ctx, ws.ToString(), filepath.Join(path...)).Result(); err != nil {
		return false, err
	} else if exists {
		if entry, err := redisClient.HGet(ctx, ws.ToString(), filepath.Join(path...)).Result(); err != nil {
			return false, err
		} else if err = json.Unmarshal([]byte(entry), &cacheEntry); err != nil {
			return false, err
		} else {
			return !cacheEntry.Deleted, nil
		}
	}

	return exists, nil
}

// GetFilesToCache saves some files of a workspace to the cache
// If an error occurs in coping one of the files clear the workspace, beacuse we want an atomic copy
func GetFilesToCache(ws datastructures.Workspace, files []datastructures.File) error {
	for _, file := range files {
		b, _ := json.Marshal(datastructures.CacheEntry{
			Data:    file.Data,
			IsDir:   file.IsDir,
			Deleted: false,
		})

		if err := redisClient.HSet(ctx, ws.ToString(), filepath.Join(file.Path...), b).Err(); err != nil {
			ClearFilesFromCache(ws)
			return err
		}
	}

	return nil
}

// GetWorkspaceToCache gets the workspace from the disk and copies it to the cache
func GetWorkspaceToCache(ws datastructures.Workspace) error {
	if files, err := GetFilesFromDisk(ws); err != nil {
		return err
	} else if err = GetFilesToCache(ws, files); err != nil {
		return err
	}

	return nil
}

// GetFilesFromCache returns the files for a workspace from the cache
func GetFilesFromCache(ws datastructures.Workspace) ([]datastructures.File, error) {
	res, err := redisClient.HGetAll(ctx, ws.ToString()).Result()

	if err != nil {
		return nil, err
	}

	var files []datastructures.File
	sep := string(os.PathSeparator)

	for name, entry := range res {
		var cacheEntry datastructures.CacheEntry

		if err = json.Unmarshal([]byte(entry), &cacheEntry); err != nil {
			return nil, err
		}

		if !cacheEntry.Deleted {
			files = append(files, datastructures.File{
				Path:  strings.Split(name, sep),
				Data:  cacheEntry.Data,
				IsDir: cacheEntry.IsDir,
			})
		}
	}

	return files, nil
}

// ClearFilesFromCache deletes all the files coresponding to a workspace from cache
func ClearFilesFromCache(ws datastructures.Workspace) error {
	return redisClient.Del(ctx, ws.ToString()).Err()
}

// CreateFile creates a new file in a workspace in cache
func CreateFile(ws datastructures.Workspace, file datastructures.File) error {
	if workspaceInCache, err := IsWorkspaceInCache(ws); err != nil {
		return err
	} else if !workspaceInCache {
		if err = GetWorkspaceToCache(ws); err != nil {
			return nil
		}
	} else if fileInCache, err := IsFileInCache(ws, file.Path); err != nil {
		return err
	} else if fileInCache {
		return errors.New("File already exists in cache")
	}

	b, _ := json.Marshal(datastructures.CacheEntry{
		Data:    file.Data,
		IsDir:   file.IsDir,
		Deleted: false,
	})

	return redisClient.HSet(ctx, ws.ToString(), filepath.Join(file.Path...), b).Err()
}

// DeleteFile marks deleted a file from a workspace in cache
// If it is a folder it marks all children as deleted
func DeleteFile(ws datastructures.Workspace, path []string) error {
	if workspaceInCache, err := IsWorkspaceInCache(ws); err != nil {
		return err
	} else if !workspaceInCache {
		if err = GetWorkspaceToCache(ws); err != nil {
			return nil
		}
	} else if fileInCache, err := IsFileInCache(ws, path); err != nil {
		return err
	} else if !fileInCache {
		return errors.New("File does not exist")
	}

	res, err := redisClient.HGetAll(ctx, ws.ToString()).Result()

	if err != nil {
		return err
	}

	pathName := filepath.Join(path...)
	b, _ := json.Marshal(datastructures.CacheEntry{Deleted: true})

	for name := range res {
		if strings.Contains(name, pathName) {
			if err = redisClient.HSet(ctx, ws.ToString(), filepath.Join(path...), b).Err(); err != nil {
				return err
			}
		}
	}

	return nil
}

// RenameFile changes the name of a file in a workspace in cache
func RenameFile(ws datastructures.Workspace, path []string, newName string) error {
	if workspaceInCache, err := IsWorkspaceInCache(ws); err != nil {
		return err
	} else if !workspaceInCache {
		if err = GetWorkspaceToCache(ws); err != nil {
			return nil
		}
	} else if fileInCache, err := IsFileInCache(ws, path); err != nil {
		return err
	} else if !fileInCache {
		// return errors.New("File not in cache")
	}

	res, err := redisClient.HGet(ctx, ws.ToString(), filepath.Join(path...)).Result()

	if err != nil {
		return err
	}

	newPath := append(append([]string(nil), path...)[:len(path)-1], newName)

	if err = redisClient.HSet(ctx, ws.ToString(), filepath.Join(newPath...), res).Err(); err != nil {
		return err
	}

	return DeleteFile(ws, path)
}

// UpdateFile applies a change to a file
func UpdateFile(ws datastructures.Workspace, path []string, change datastructures.Change) error {
	if workspaceInCache, err := IsWorkspaceInCache(ws); err != nil {
		return err
	} else if !workspaceInCache {
		if err = GetWorkspaceToCache(ws); err != nil {
			return nil
		}
	} else if fileInCache, err := IsFileInCache(ws, path); err != nil {
		return err
	} else if !fileInCache {
		return errors.New("File does not exist")
	}

	res, err := redisClient.HGet(ctx, ws.ToString(), filepath.Join(path...)).Result()

	if err != nil {
		return err
	}

	var cacheEntry datastructures.CacheEntry
	if err = json.Unmarshal([]byte(res), &cacheEntry); err != nil {
		return err
	}

	if cacheEntry.IsDir {
		return errors.New("Folder has no data to modify")
	} else if cacheEntry.Deleted {
		return errors.New("File has been deleted")
	}

	data := []byte(cacheEntry.Data)

	var last []byte

	start := change.Position
	end := change.Position + int64(len(change.Previous))

	if int64(len(data)) >= end {
		last = make([]byte, len(data[end:]))
		copy(last, data[end:])
	}

	cacheEntry.Data = string(append(append(data[:start], []byte(change.Current)...), last...))

	b, _ := json.Marshal(cacheEntry)

	return redisClient.HSet(ctx, ws.ToString(), filepath.Join(path...), b).Err()
}
