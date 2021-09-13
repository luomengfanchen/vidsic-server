package model

import "time"

// 文章数据结构体
type Article struct {
	Id     int       `json:"id"`
	Name   string    `json:"name"`
	Date   time.Time `json:"date"`
	Author int       `json:"author"`
	Cover  string    `json:"cover"`
	Text   string    `json:"text"`
	Views  int       `json:"views"`
}
