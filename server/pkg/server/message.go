package server

type ClientCommand struct {
	Action   string `json:"action"`
	SrcID    string `json:"srcId"`
	TargetID string `json:"targetId"`
}
