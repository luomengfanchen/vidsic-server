package model

import "fmt"

// 查询
func QueryRowCommit(id int, name string, date string) (int, error) {
	var respId int

	// 预加载sql
	stmt, err := Db.Prepare("SELECT id FROM commit_t WHERE commiter = ? AND name = ? AND date = ?")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()

	// 查询数据
	err = stmt.QueryRow(id, name, date).Scan(&respId)
	if err != nil {
		fmt.Println(err.Error())
	}

	return respId, err
}