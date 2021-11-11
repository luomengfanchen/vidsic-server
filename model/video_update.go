package model

import "fmt"

// 更新视频数据，将播放量+1
func UpdateVideo(video Video) error {
	sql := "UPDATE video_t SET views = views + 1 WHERE id = ?"

	stmt, err := Db.Prepare(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(video.Id)
	if err != nil {
		fmt.Println(err.Error())
	}

	return err
}
