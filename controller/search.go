package controller

import (
	"net/http"
	"vidsic/model"

	"github.com/gin-gonic/gin"
)

// 处理搜索请求,返回json数组
func Search(c *gin.Context) {
	// 获取搜索类型
	searchType := c.Query("type")
	if searchType == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
		return
	}

	// 获取搜索类型
	searchVal := c.Query("q")
	if searchVal == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
		return
	}

	if searchType == "video" {
		info, err := model.QueryVideoOfName(searchVal)
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
	} else if searchType == "music" {
		info, err := model.QueryMusicOfName(searchVal)
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
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
	}
}
