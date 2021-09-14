package model

import "fmt"

func InsertCommit(commit Commit) error {
	// sql
	sql := "INSERT INTO commit_t (name, date, path, type, commiter) VALUES (?, ?, ?, ?, ?)"

	// 预加载sql语句
	stmt, err := Db.Prepare(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()

	// 执行sql语句
	_, err = stmt.Exec(commit.Name, commit.Date, commit.Path, commit.Type, commit.Commiter)
	if err != nil {
		fmt.Println(err.Error())
	}

	return err
}
