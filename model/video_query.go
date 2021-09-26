package model

import "fmt"

// 查询最新的视频数据
func QueryVideo(limit int64) ([]Video, error) {
	// 临时保存行数据
	var video Video
	// 最近视频列表
	var videoList []Video

	// 按照上传时间降序排列，并限制数据条数
	sql := "SELECT id, name, date, author, cover, type, path, views FROM video_t ORDER BY date DESC LIMIT ?"

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
		err := rows.Scan(&video.Id, &video.Name, &video.Date, &video.Author, &video.Cover, &video.Type, &video.Path, &video.Views)
		if err != nil {
			fmt.Println(err.Error())
		}

		// 将单个条目添加到数组中
		videoList = append(videoList, video)
	}

	return videoList, err
}

// 通过id查询指定的视频
func QueryRowVideo(id int) (Video, error){
	sql := "SELECT id, name, descript, date, author, cover, type, path, views FROM video_t WHERE id = ?"
	stmt, err := Db.Prepare(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()

	var video Video
	err = stmt.QueryRow(id).Scan(&video.Id, &video.Name, &video.Descript, &video.Date, &video.Author, &video.Cover, &video.Type, &video.Path, &video.Views)
	if err != nil {
		fmt.Println(err.Error())
	}

	return video, err
}

// 通过name查询指定的视频
func QueryVideoOfName(name string) ([]Video, error){
	// 临时保存行数据
	var video Video
	// 最近视频列表
	var videoList []Video

	// 通过name进行模糊搜索
	sql := "SELECT id, name, date, author, cover, type, path, views FROM video_t WHERE name LIKE ?"

	// 预加载sql语句
	stmt, err := Db.Prepare(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()

	// 执行sql语句
	rows, err := stmt.Query("%" + name + "%")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	// 提取数据
	for rows.Next() {
		err := rows.Scan(&video.Id, &video.Name, &video.Date, &video.Author, &video.Cover, &video.Type, &video.Path, &video.Views)
		if err != nil {
			fmt.Println(err.Error())
		}

		// 将单个条目添加到数组中
		videoList = append(videoList, video)
	}

	return videoList, err
}
