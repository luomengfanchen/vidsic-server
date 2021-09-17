package model

import "fmt"

func InsertMusic(info Music) error {
	sql := "INSERT INTO music_t (name, date, singer, descript, cover, type, path) VALUES (?, ?, ?, ?, ?, ?, ?)"
	stmt, err := Db.Prepare(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(info.Name, info.Date, info.Singer, info.Descript, info.Cover, info.Type, info.Path)
	if err != nil {
		fmt.Println(err.Error())
	}

	return err
}