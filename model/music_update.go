package model

import "fmt"

func UpdateMusic(music Music) error {
	sql := "UPDATE music_t SET views = views + 1 WHERE id = ?"

	stmt, err := Db.Prepare(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(music.Id)
	if err != nil {
		fmt.Println(err.Error())
	}

	return err
}
