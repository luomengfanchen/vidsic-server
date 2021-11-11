package model

import (
	"fmt"
)

// 添加新用户到user表中
func InsertUser(user User) error {
	// 插入用户sql语句
	sql := "INSERT INTO user_t (nickname, email, password, birthday, intro, token)"
	sql += "VALUES (?, ?, ?, ?, ?, ?)"

	// 预加载sql语句
	stmt, err := Db.Prepare(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()

	// 执行sql语句
	_, err = stmt.Exec(user.NickName, user.Email, user.Password, user.Birthday, user.Intro, user.Token)
	if err != nil {
		fmt.Println(err.Error())
	}

	return err
}
