package datastructures

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
	File      struct {
		Path  []string `json:"path"`
		Data  string   `json:"data"`
		IsDir bool     `json:"isdir"`
	} `json:"file"`
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
