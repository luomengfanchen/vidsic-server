package utils

import "vidsic/model"

// 响应数据结构体
type ResponseVideo struct {
	Id       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Descript string `json:"descript,omitempty"`
	Date     string `json:"date,omitempty"`
	Author   string `json:"author,omitempty"`
	Cover    string `json:"cover,omitempty"`
	Type     string `json:"type,omitempty"`
	Path     string `json:"path,omitempty"`
	Views    int    `json:"views,omitempty"`
}

// 转换视频响应数据
func ConverVIdeoResponse(videoList []model.Video) ([]ResponseVideo, error) {
	var list []ResponseVideo

	for _, ele := range videoList {
		user, err := model.QueryRowUserOfId(ele.Author)
		if err != nil {
			return []ResponseVideo{}, err
		}

		item := ResponseVideo{
			Id: ele.Id,
			Name: ele.Name,
			Descript: ele.Descript,
			Date: ele.Date,
			Author: user.NickName,
			Cover: ele.Cover,
			Type: ele.Type,
			Views: ele.Views,
		}

		list = append(list, item)
	}

	return list, nil
}
