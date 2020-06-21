package datastructures

import (
	"path/filepath"
	"strconv"
)

const fileSystem = "/data"

// Workspace is a struct that defines the folder for an assignment
// It is used for getting the key from cache
type Workspace struct {
	Owner          string   `json:"owner"`
	Year           int      `json:"year"`
	Course         string   `json:"course"`
	Series         string   `json:"series"`
	AssignmentName string   `json:"assignmentname"`
	Folder         []string `json:"folder"`
}

// ToString transforms the workspace struct into a path
func (ws Workspace) ToString() string {
	return filepath.Join(fileSystem, strconv.Itoa(ws.Year), ws.Course, ws.Series, ws.AssignmentName, ws.Owner, filepath.Join(ws.Folder...))
}

// File struct contains the path to the file, the data, and whether or not the file is a folder
// It is used for transmitting files
type File struct {
	Path  []string `json:"path"`
	Data  string   `json:"data"`
	IsDir bool     `json:"isdir"`
}

// CacheEntry is the structure of the data stored in cache
// It is used for working with the cache
type CacheEntry struct {
	Data    string
	IsDir   bool
	Deleted bool
}

// Change is the struct for changes of one file
// It contains the position of the change, what we desire to have at that position and what was previously at that position
type Change struct {
	Position int64  `json:"position"`
	Current  string `json:"current"`
	Previous string `json:"previous"`
}

// CreateBody is the format for the create request body
type CreateBody struct {
	Workspace Workspace `json:"workspace"`
	File      File      `json:"file"`
}

// DeleteBody is the format for the delete request body
type DeleteBody struct {
	Workspace Workspace `json:"workspace"`
	Path      []string  `json:"path"`
}

// RenameBody is the format for the rename request body
type RenameBody struct {
	Workspace Workspace `json:"workspace"`
	Path      []string  `json:"path"`
	NewName   string    `json:"newname"`
}

// UpdateBody is the format for the update request body
type UpdateBody struct {
	Sender    string    `json:"sender"`
	Workspace Workspace `json:"workspace"`
	Path      []string  `json:"path"`
	Change    Change    `json:"change"`
}
