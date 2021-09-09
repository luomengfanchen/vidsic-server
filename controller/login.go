package controller

import (
	"fmt"
	"net/http"
	"vidsic/model"
	"vidsic/utils"

	"github.com/gin-gonic/gin"
)

// 处理登录请求
func Login(c *gin.Context) {
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

	// 处理数据，加密密码
	user.Password = utils.EncodeMd5(user.Password)

	// 数据库操作，查询用户
	user, err := model.QueryRowUserOfInfo(user.Email, user.Password)
	if err != nil {
		fmt.Println(err.Error())

		// 返回错误响应
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
		return
	}

	// 返回成功响应数据
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
		"data": user,
	})
}
