package model

import "fmt"

// 向视频表插入新的数据
func InsertVideo(video Video) error {
	// sql
	sql := "INSERT INTO video_t (name, date, author, cover, type, path) VALUES (?, ?, ?, ?, ?, ?)"

	// 预加载sql语句
	stmt, err := Db.Prepare(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()

	// 执行sql语句
	_, err = stmt.Exec(video.Name, video.Date, video.Author, video.Cover, video.Type, video.Path)
	if err != nil {
		fmt.Println(err.Error())
	}

	return err
}
