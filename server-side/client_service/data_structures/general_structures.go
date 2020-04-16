package datastructures

import (
	"path/filepath"
	"strconv"
)

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
	return filepath.Join(strconv.Itoa(ws.Year), ws.Course, ws.Series, ws.AssignmentName, ws.Owner, filepath.Join(ws.Folder...))
}
