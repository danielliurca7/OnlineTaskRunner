package tree

type Node struct {
	Name     string   `json:"name"`
	Children []*Node  `json:"children"`
	IsDir    bool     `json:"isDir"`
	Path     []string `json:"path"`
}
