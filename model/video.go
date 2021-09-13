package model

// 视频数据结构体
type Video struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Date   string `json:"date"`
	Author int    `json:"author"`
	Cover  string `json:"cover"`
	Type   string `json:"type"`
	Path   string `json:"path"`
	Views  int    `json:"views"`
}
