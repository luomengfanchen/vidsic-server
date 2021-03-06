package main

import (
	"net/http"
	"vidsic/controller"
	"vidsic/middleware"
	"vidsic/utils"

	"github.com/gin-gonic/gin"
)

// 加载配置文件
func init() {
	utils.LoadConfig("config.json")
}

// 路由监听配置
func main() {
	r := gin.Default()

	// 页面服务
	r.LoadHTMLFiles(utils.Config.VueDist + "/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.Static("/static", utils.Config.VueDist + "/static")

	// api服务
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

		// 获取指定视频
		api.GET("/video/:id", controller.GetVideo)

		// 获取最新音乐
		api.GET("/music", controller.LatestMusic)

		// 获取指定音乐
		api.GET("/music/:id", controller.GetMusic)

		// 搜索接口
		api.GET("/search", controller.Search)

		// 使用验证中间件验证用户身份
		auth := api.Group("/auth",middleware.AuthorizeMiddle)
		{
			// 获取用户信息
			auth.GET("/info", controller.GetInfo)

			// 修改用户信息
			auth.PATCH("/info", controller.UpdateInfo)

			// 接收上传文件
			auth.POST("/commit", controller.Commit)

			// 接收上传文件信息
			auth.POST("/upload", controller.Upload)
		}
	}

	r.Run(utils.Config.Port)
}
