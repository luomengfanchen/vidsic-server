package model

import "fmt"

// 查询最新的音乐数据
func QueryMusic(limit int64) ([]Music, error) {
	// 临时保存行数据
	var music Music
	// 最近视频列表
	var musicList []Music

	// 按照上传时间降序排列，并限制数据条数
	sql := "SELECT id, name, date, singer, cover, type, path, views FROM music_t ORDER BY date DESC LIMIT ?"

	// 预加载sql语句
	stmt, err := Db.Prepare(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()

	// 执行sql语句
	rows, err := stmt.Query(limit)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	// 提取数据
	for rows.Next() {
		err := rows.Scan(&music.Id, &music.Name, &music.Date, &music.Singer, &music.Cover, &music.Type, &music.Path, &music.Views)
		if err != nil {
			fmt.Println(err.Error())
		}

		// 将单个条目添加到数组中
		musicList = append(musicList, music)
	}

	return musicList, err
}

// 通过id查询指定的视频
func QueryRowMusic(id int) (Music, error){
	sql := "SELECT id, name, date, singer, cover, type, path, views FROM music_t WHERE id = ?"
	stmt, err := Db.Prepare(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()

	var music Music
	err = stmt.QueryRow(id).Scan(&music.Id, &music.Name, &music.Date, &music.Singer, &music.Cover, &music.Type, &music.Path, &music.Views)
	if err != nil {
		fmt.Println(err.Error())
	}

	return music, err
}

// 通过name查询指定英语
func QueryMusicOfName(name string) ([]Music, error) {
	// 临时保存行数据
	var music Music
	// 最近视频列表
	var musicList []Music

	// 按照上传时间降序排列，并限制数据条数
	sql := "SELECT id, name, date, singer, cover, type, path, views FROM music_t WHERE name LIKE '%' || ? || '%'"

	// 预加载sql语句
	stmt, err := Db.Prepare(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()

	// 执行sql语句
	rows, err := stmt.Query(name)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	// 提取数据
	for rows.Next() {
		err := rows.Scan(&music.Id, &music.Name, &music.Date, &music.Singer, &music.Cover, &music.Type, &music.Path, &music.Views)
		if err != nil {
			fmt.Println(err.Error())
		}

		// 将单个条目添加到数组中
		musicList = append(musicList, music)
	}

	return musicList, err
}