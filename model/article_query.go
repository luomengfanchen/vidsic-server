package model

import "fmt"

// 查询最新的文章数据
func QueryArticle(limit int64) ([]Article, error) {
	// 临时保存行数据
	var article Article
	// 最近视频列表
	var articleList []Article

	// 按照上传时间降序排列，并限制数据条数
	sql := "SELECT id, name, date, author, cover, text, views FROM article_t ORDER BY date DESC LIMIT ?"

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
		err := rows.Scan(&article.Id, &article.Name, &article.Date, &article.Author, &article.Cover, &article.Text, &article.Views)
		if err != nil {
			fmt.Println(err.Error())
		}

		// 将单个条目添加到数组中
		articleList = append(articleList, article)
	}

	return articleList, err
}
