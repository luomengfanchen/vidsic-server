package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"vidsic/model"
	"vidsic/utils"

	"github.com/gin-gonic/gin"
)

// 上传文件
func Commit(c *gin.Context) {
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

	// 接收上传文件
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
		return
	}

	// 设置保存文件的路径
	nowTime := time.Now().Unix()
	timeStr := strconv.Itoa(int(nowTime))

	baseURL := utils.Config.StaticPath
	if filetype == "video" {
		baseURL += "/video/"
	} else if baseURL == "music" {
		baseURL += "/music/"
	} else {
		baseURL += "/image/"
	}

	// 保存文件
	err = c.SaveUploadedFile(file, baseURL+timeStr+"@"+file.Filename)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
		return
	}

	// 生成表数据
	commit := model.Commit{
		Name:     timeStr + "@" + file.Filename,
		Date:     time.Now().Format("2006-01-02 15:04:05"),
		Type:     filetype,
		Path:     baseURL + file.Filename,
		Commiter: id,
	}

	// 数据库操作，将数据插入数据表
	err = model.InsertCommit(commit)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
		return
	}

	// 查询响应的数据
	commitId, err := model.QueryRowCommit(commit.Commiter, commit.Name, commit.Date)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
		"date": gin.H{
			"id": commitId,
		},
	})
}
