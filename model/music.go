package model

// 音乐数据结构体
type Music struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Date   string `json:"date"`
	Singer string `json:"singer"`
	Cover  string `json:"cover"`
	Type   string `json:"type"`
	Path   string `json:"path"`
	Views  int    `json:"views"`
}
