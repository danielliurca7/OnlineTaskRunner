package datastructures

import (
	"fmt"
	"path/filepath"
	"strconv"
)

const tmp = "/tmp"

var dockerfile = func(image string) string {
	return fmt.Sprintf(`FROM %s
WORKDIR /app
COPY . .
RUN apt update
RUN ./build.sh
CMD ./run.sh`, image)
}

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
	return filepath.Join(tmp, strconv.Itoa(ws.Year), ws.Course, ws.Series, ws.AssignmentName, ws.Owner, filepath.Join(ws.Folder...))
}

// Name transforms the workspace struct into a path, but excluding owner
func (ws Workspace) Name() string {
	return filepath.Join(tmp, strconv.Itoa(ws.Year), ws.Course, ws.Series, ws.AssignmentName)
}

// BuildBody is the body structure for build requests
type BuildBody struct {
	Image     string    `json:"image"`
	Workspace Workspace `json:"workspace"`
}

// File struct contains the path to the file, the data, and whether or not the file is a folder
// It is used for transmitting files
type File struct {
	Path  []string `json:"path"`
	Data  string   `json:"data"`
	IsDir bool     `json:"isdir"`
}

// GetDockerFile build the dockerfile text starting from the image given as parameter
func GetDockerFile(image string) File {
	return File{
		Path:  []string{"Dockerfile"},
		Data:  dockerfile(image),
		IsDir: false,
	}
}
