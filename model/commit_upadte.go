package model

import "fmt"

// 更新提交状态
func UpadteCommit(id int) error {
	stmt, err := Db.Prepare("UPDATE commit_t SET status = TRUE WHERE id = ?")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		fmt.Println(err.Error())
	}

	return err
}
