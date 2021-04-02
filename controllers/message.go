package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhuziqi1999/shuwo/models"
	"net/http"
)

func GetLikeMessageList(c *gin.Context) {
	var (
		likelist interface{}
		user     = &models.User{}
	)

	if err := c.ShouldBindJSON(&user); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("err:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
		})
		return
	}

	likelist = models.GetLikeMessageList(user.UserOpenid)

	c.JSON(http.StatusOK, gin.H{
		"code":     1,
		"likelist": likelist,
	})
}

func GetCommentMessageList(c *gin.Context) {
	var (
		commentlist interface{}
		user        = &models.User{}
	)

	if err := c.ShouldBindJSON(&user); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("err:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
		})
		return
	}

	commentlist = models.GetCommentMessageList(user.UserOpenid)

	c.JSON(http.StatusOK, gin.H{
		"code":        1,
		"commentlist": commentlist,
	})
}

func GetCollectMessageList(c *gin.Context) {
	var (
		collectlist interface{}
		user        = &models.User{}
	)

	if err := c.ShouldBindJSON(&user); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("err:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
		})
		return
	}

	collectlist = models.GetCollectMessageList(user.UserOpenid)

	c.JSON(http.StatusOK, gin.H{
		"code":        1,
		"collectlist": collectlist,
	})
}
