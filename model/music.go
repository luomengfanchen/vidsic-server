package model

import "time"

// 音乐数据结构体
type Music struct {
	Id     int       `json:"id"`
	Name   string    `json:"name"`
	Date   time.Time `json:"date"`
	Singer string    `json:"singer"`
	Cover  string    `json:"cover"`
	Type   string    `json:"type"`
	Path   string    `json:"path"`
	Views  int       `json:"views"`
}
