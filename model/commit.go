package model

type Commit struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Date     string `json:"date"`
	Type     string `json:"type"`
	Path     string `json:"path"`
	Status   bool   `json:"status"`
	Commiter int    `json:"commit"`
}
