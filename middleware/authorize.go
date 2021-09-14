package middleware

import (
	"net/http"
	"vidsic/model"

	"github.com/gin-gonic/gin"
)

// 用户验证中间件
func AuthorizeMiddle(c *gin.Context) {
	// 获取http头部信息
	authorization := c.Request.Header.Get("Authorization")

	// 验证获取的token
	user, err := model.QueryRowUserOfToken(authorization)

	// 验证失败时终止后续执行操作
	if err != nil {
		// 返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
		// 终止后续执行
		c.Abort()
	} else {
		// 设置需要后续使用的数据
		c.Set("id", user.Id)
		c.Set("name", user.NickName)
		c.Set("comptence", user.Competence)

		// 继续执行后续操作
		c.Next()
	}
}
