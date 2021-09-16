package model

// 音乐数据结构体
type Music struct {
	Id       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Descript string `json:"descript,omitempty"`
	Date     string `json:"date,omitempty"`
	Singer   string `json:"singer,omitempty"`
	Cover    string `json:"cover,omitempty"`
	Type     string `json:"type,omitempty"`
	Path     string `json:"path,omitempty"`
	Views    int    `json:"views,omitempty"`
}
