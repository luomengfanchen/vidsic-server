package main

import (
	"vidsic/controller"
	"vidsic/middleware"
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

		// 获取最新视频
		api.GET("/video", controller.LatestVideo)

		// 获取最新文章
		api.GET("/article", controller.LatestArticle)

		// 获取最新音乐
		api.GET("/music", controller.LatestMusic)

		// 使用验证中间件验证用户身份
		auth := api.Group("/auth",middleware.AuthorizeMiddle)
		{
			auth.POST("/commit", controller.Commit)
		}
	}

	r.Run(utils.Config.Port)
}
