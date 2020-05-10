package datastructures

// SubscribeBody is the format for the subscribe request body
type SubscribeBody struct {
	Token     string    `json:"token"`
	Workspace Workspace `json:"workspace"`
}

// ChangeBody is the format for the change request body
type ChangeBody struct {
	Token  string     `json:"token"`
	Sender string     `json:"sender"`
	Update UpdateBody `json:"update"`
}
