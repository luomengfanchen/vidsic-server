package controller

import (
	"net/http"
	"vidsic/model"
	"vidsic/utils"

	"github.com/gin-gonic/gin"
)

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
		var videoInfo model.Video
		c.BindJSON(&videoInfo)

		videoInfo.Author = id
		videoInfo.Date = utils.NowTime()

		// 数据库操作
		err := model.InsertVideo(videoInfo)
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
		var musicInfo model.Music
		c.BindJSON(&musicInfo)

		musicInfo.Date = utils.NowTime()

		// 数据库操作
		err := model.InsertMusic(musicInfo)
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

		// 当上传的为图片时
	} else if filetype == "image" {
		// 接收json
		var imageInfo model.Image
		c.BindJSON(&imageInfo)

		imageInfo.Author = id
		imageInfo.Date = utils.NowTime()

		// 数据库操作
		err := model.InsertImage(imageInfo)
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
