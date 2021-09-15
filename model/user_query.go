package model

import "fmt"

// 通过用户id查询用户信息
func QueryRowUserOfInfo(email string, password string) (User, error) {
	user := User{}

	// 预加载sql
	stmt, err := Db.Prepare("SELECT nickname, avator, token FROM user_t WHERE email = ? AND password = ?")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()

	// 查询数据
	err = stmt.QueryRow(email, password).Scan(&user.NickName, &user.Avator, &user.Token)
	if err != nil {
		fmt.Println(err.Error())
	}

	return user, err
}

// 通过用户token查询用户信息
func QueryRowUserOfToken(token string) (User, error) {
	user := User{}

	// 预加载sql
	stmt, err := Db.Prepare("SELECT id, nickname, competence FROM user_t WHERE token = ?")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()

	// 查询数据
	err = stmt.QueryRow(token).Scan(&user.Id, &user.NickName, &user.Competence)
	if err != nil {
		fmt.Println(err.Error())
	}

	return user, err
}
