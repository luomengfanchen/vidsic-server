package controller

import (
	"net/http"
	"vidsic/model"
	"vidsic/utils"

	"github.com/gin-gonic/gin"
)

type recvUpload struct {
	Name     string `json:"name,omitempty"`
	Descript string `json:"descript,omitempty"`
	Singer   string `json:"singer,omitempty"`
	Type     string `json:"type,omitempty"`
	Media    int    `json:"media,omitempty"`
	Cover    int    `json:"cover,omitempty"`
}

// 接收上传文件信息
func Upload(c *gin.Context) {
	// 获取上传者id
	tmp, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
		return
	}
	id := tmp.(int)

	// 获取文件类型
	filetype := c.Query("type")
	if filetype == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
		return
	}

	// 当上传的为视频时
	if filetype == "video" {
		// 接收json
		var recvInfo recvUpload
		c.BindJSON(&recvInfo)

		// 查询视频路径
		videoPath, err := model.QueryRowCommitOfId(recvInfo.Media)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "error",
			})
			return
		}

		// 查询封面路径
		coverPath, err := model.QueryRowCommitOfId(recvInfo.Cover)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "error",
			})
			return
		}

		// 生成插入的数据
		videoInfo := model.Video{
			Name:     recvInfo.Name,
			Descript: recvInfo.Descript,
			Type:     recvInfo.Type,
			Author:   id,
			Date:     utils.NowTime(),
			Path:     videoPath,
			Cover:    coverPath,
		}

		// 数据库操作
		err = model.InsertVideo(videoInfo)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "error",
			})
			return
		}

		// 修改提交状态为True
		err = model.UpadteCommit(recvInfo.Media)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "error",
			})
			return
		}
		err = model.UpadteCommit(recvInfo.Cover)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "error",
			})
			return
		}

		// 返回响应
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
		// 当上传的为音乐时
	} else if filetype == "music" {
		// 接收json
		var recvInfo recvUpload
		c.BindJSON(&recvInfo)

		// 查询视频路径
		musicPath, err := model.QueryRowCommitOfId(recvInfo.Media)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "error",
			})
			return
		}

		// 查询封面路径
		coverPath, err := model.QueryRowCommitOfId(recvInfo.Cover)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "error",
			})
			return
		}

		musicInfo := model.Music{
			Name:     recvInfo.Name,
			Descript: recvInfo.Descript,
			Singer: recvInfo.Singer,
			Type:     recvInfo.Type,
			Date:     utils.NowTime(),
			Path:     musicPath,
			Cover:    coverPath,
		}

		// 数据库操作
		err = model.InsertMusic(musicInfo)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "error",
			})
			return
		}

		// 返回响应
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
	}
}
