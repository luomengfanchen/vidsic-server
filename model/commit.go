package model

type Commit struct {
	Id       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Date     string `json:"date,omitempty"`
	Type     string `json:"type,omitempty"`
	Path     string `json:"path,omitempty"`
	Status   bool   `json:"status,omitempty"`
	Commiter int    `json:"commit,omitempty"`
}
