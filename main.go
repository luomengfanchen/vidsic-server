package main

import (
	"vidsic/controller"
	"vidsic/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	utils.LoadConfig("config.json")
}

func main() {
	r := gin.Default()

	api := r.Group("/api")
	{
		// api静态文件服务
		api.Static("/static", utils.Config.StaticPath)

		// 登录接口
		api.POST("/login", controller.Login)

		// 注册接口
		api.POST("/register", controller.Register)
	}

	r.Run(utils.Config.Port)
}