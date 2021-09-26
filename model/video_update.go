package model

import "fmt"

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
