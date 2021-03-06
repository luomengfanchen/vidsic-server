package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"vidsic/model"

	"github.com/gin-gonic/gin"
)

// 获取最新的音乐列表
func LatestMusic(c *gin.Context) {
	if num := c.Query("num"); num != "" {
		// 类型转换
		limit, err := strconv.ParseInt(num, 10, 32)
		if err != nil {
			fmt.Println(err.Error())
			// 返回错误响应
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "error",
			})
			return
		}

		// 查询数据
		list, err := model.QueryMusic(limit)
		if err != nil {
			fmt.Println(err.Error())
			// 返回错误响应
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "error",
			})
			return
		}

		// 返回成功响应
		c.JSON(http.StatusOK, gin.H{
			"msg":  "ok",
			"data": list,
		})
	} else {
		// 返回错误响应
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
	}
}

// 获取指定的音乐信息
func GetMusic(c *gin.Context) {
	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		// 返回错误响应
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
		return
	}

	info, err := model.QueryRowMusic(id)
	if err != nil {
		// 返回错误响应
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
		return
	}

	// 更改数据库，播放量+1
	err = model.UpdateMusic(info)
	if err != nil {
		// 返回错误响应
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
		"data": info,
	})
}
