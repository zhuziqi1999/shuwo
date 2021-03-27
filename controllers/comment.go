package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhuziqi1999/shuwo/models"
	"log"
	"net/http"
)

func CreateComment(c *gin.Context) {
	var (
		comment = &models.Comment{}
		res     = gin.H{}
	)

	if err := c.ShouldBindJSON(&comment); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("err:", err)
		return
	}

	comment, err := models.CreateComment(comment.CommentCreatedBy, comment.CommentText, comment.CommentContentID, comment.CommentBackUserID)

	if err != nil {
		log.Println(err)
		res["code"] = 0
		res["message"] = err.Error()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"user": map[string]interface{}{
			"ContentID":       comment.CommentID,
			"ContentCreateBy": comment.CommentCreatedBy,
		},
	})

}

func GetCommentList(c *gin.Context) {
	var (
		comment interface{}
		content = &models.Content{}
	)

	if err := c.ShouldBindJSON(&content); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("err:", err)
		return
	}

	comment = models.GetCommentList(content.ContentID)

	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"comment": comment,
	})

}
