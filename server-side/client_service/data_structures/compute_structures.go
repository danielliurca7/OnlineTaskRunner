package datastructures

// BuildBody is the body structure for build requests
type BuildBody struct {
	Image     string    `json:"image"`
	Workspace Workspace `json:"workspace"`
}

// ExecBody is the body structure for exec requests
type ExecBody struct {
	Command   string    `json:"command"`
	Workspace Workspace `json:"workspace"`
}
