package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhuziqi1999/shuwo/models"
	"net/http"
)

func SearchContent(c *gin.Context) {

	var (
		content interface{}
		search  models.Search
	)

	if err := c.ShouldBindJSON(&search); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("err:", err)
		return
	}

	content = models.SearchContent(search.UserOpenID, search.Message)

	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"content": content,
	})

}

func SearchGroup(c *gin.Context) {
	var (
		group  interface{}
		search models.Search
	)

	if err := c.ShouldBindJSON(&search); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("err:", err)
		return
	}

	group = models.SearchGroup(search.UserOpenID, search.Message)

	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"content": group,
	})

}

func SearchFile(c *gin.Context) {
	var (
		search   models.Search
		filelist interface{}
	)

	if err := c.ShouldBindJSON(&search); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("err:", err)
		return
	}

	filelist = models.SearchFile(search.UserOpenID, search.Message)

	c.JSON(http.StatusOK, gin.H{
		"code":     1,
		"filelist": filelist,
	})

}
