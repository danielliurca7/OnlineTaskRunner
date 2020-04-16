package datastructures

// BuildBody is the body structure for build requests
type BuildBody struct {
	Image     string    `json:"image"`
	Workspace Workspace `json:"workspace"`
}
