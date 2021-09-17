package controller

import (
	"net/http"
	"vidsic/model"

	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {
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
			"msg": "ok",
			"data": info,
		})
	}
}
