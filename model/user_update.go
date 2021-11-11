package model

import "fmt"

// 更新用户信息
func UpdateUser(user User) error {
	sql := "UPDATE user_t SET email = ?, nickname = ?, avator = ?, birthday = ?, intro = ? WHERE id = ?"
	stmt, err := Db.Prepare(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Email, user.NickName, user.Avator, user.Birthday, user.Intro, user.Id)
	if err != nil {
		fmt.Println(err.Error())
	}

	return err
}
