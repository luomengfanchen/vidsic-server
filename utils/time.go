package utils

import "time"

// 返回现在的时间
func NowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}