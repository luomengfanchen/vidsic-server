package model

// 视频数据结构体
type Video struct {
	Id       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Descript string `json:"descript,omitempty"`
	Date     string `json:"date,omitempty"`
	Author   int    `json:"author,omitempty"`
	Cover    string `json:"cover,omitempty"`
	Type     string `json:"type,omitempty"`
	Path     string `json:"path,omitempty"`
	Views    int    `json:"views,omitempty"`
}
