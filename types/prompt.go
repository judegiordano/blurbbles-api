package types

type Prompt struct {
	Id     string `json:"id"`
	Prompt string `json:"prompt"`
	Genre  Genre  `json:"genre"`
}
