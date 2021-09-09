package controller

import (
	"fmt"
	"net/http"
	"vidsic/model"
	"vidsic/utils"

	"github.com/gin-gonic/gin"
)

// 处理注册请求
func Register(c *gin.Context) {
	// 用户结构体
	var user model.User

	// 接收json数据
	if err := c.BindJSON(&user); err != nil {
		fmt.Println(err.Error())

		// 返回错误响应
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
		return
	}

	// 处理数据，加密密码，生成token
	user.Password = utils.EncodeMd5(user.Password)
	user.Token = utils.EncodeToken()

	// 向数据库添加数据
	if err := model.InsertUser(user); err != nil {
		fmt.Println(err.Error())

		// 返回错误响应
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
