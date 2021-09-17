package controller

import (
	"net/http"
	"vidsic/model"

	"github.com/gin-gonic/gin"
)

func GetInfo(c *gin.Context) {
	// 获取上传者id
	tmp, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
		return
	}
	id := tmp.(int)

	info, err := model.QueryRowUserOfId(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "ok",
			"data": info,
		})
	}
}

func UpdateInfo(c *gin.Context) {
	var user model.User
	// 获取上传者id
	tmp, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
		return
	}
	id := tmp.(int)

	// 接收json
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
		return
	} else {
		user.Id = id
	}

	// 数据库操作
	err = model.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
