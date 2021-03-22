package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"shuwo/models"
)

func CreateContent(c *gin.Context) {
	var (
		content = &models.Content{}
		res     = gin.H{}
	)

	if err := c.ShouldBindJSON(&content); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("err:", err)
		return
	}

	content, err := models.CreateContent(content.ContentCreatedBy, content.ContentText, content.ContentShare)

	if err != nil {
		log.Println(err)
		res["message"] = err.Error()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"user": map[string]interface{}{
			"ContentID":       content.ContentID,
			"ContentCreateBy": content.ContentCreatedBy,
		},
	})

}

func GetHotContentList(c *gin.Context) {
	var (
		content interface{}
	)
	content = models.GetHotContentList()

	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"content": content,
	})

}
