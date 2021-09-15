package model

import "fmt"

func InsertImage(info Image) error {
	sql := "INSERT INTO image_t (name, date, author, path) VALUES (?, ?, ?, ?)"
	stmt, err := Db.Prepare(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(info.Name, info.Date, info.Author, info.Path)
	if err != nil {
		fmt.Println(err.Error())
	}

	return err
}