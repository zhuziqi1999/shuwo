package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhuziqi1999/shuwo/models"
	"log"
	"net/http"
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

	content, err := models.CreateContent(content.ContentCreatedBy, content.ContentText, content.ContentShare, content.ContentFrom)

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

func GetContentList(c *gin.Context) {
	type n struct {
		UserOpenID string `json:"useropenid"  gorm:"column:USER_OPEN_ID"`
		GroupID    string `json:"groupid"  gorm:"column:GROUP_ID"`
	}

	var (
		content interface{}
		group   n
	)

	if err := c.ShouldBindJSON(&group); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("err:", err)
		return
	}

	content = models.GetContentList(group.UserOpenID, group.GroupID)

	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"content": content,
	})

}

func LikeContent(c *gin.Context) {
	var (
		userlikecontent models.UserLikeContent
	)

	if err := c.ShouldBindJSON(&userlikecontent); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("err:", err)
		return
	}

	err := models.LikeContent(userlikecontent.UserID, userlikecontent.ContentID)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code":    0,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "点赞成功",
	})

	return
}

func UnlikeContent(c *gin.Context) {
	var (
		userlikecontent models.UserLikeContent
	)

	if err := c.ShouldBindJSON(&userlikecontent); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("err:", err)
		return
	}

	err := models.UnlikeContent(userlikecontent.UserID, userlikecontent.ContentID)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code":    0,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "取消点赞成功",
	})

	return
}

func GetLikeList(c *gin.Context) {
	var (
		likelist interface{}
		content  = &models.Content{}
	)

	if err := c.ShouldBindJSON(&content); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("err:", err)
		return
	}

	likelist = models.GetLikeList(content.ContentID)

	c.JSON(http.StatusOK, gin.H{
		"code":     1,
		"likelist": likelist,
	})

}

func CollectContent(c *gin.Context) {
	var (
		usercollectcontent models.UserCollectContent
	)

	if err := c.ShouldBindJSON(&usercollectcontent); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("err:", err)
		return
	}

	err := models.CollectContent(usercollectcontent.UserID, usercollectcontent.ContentID)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code":    0,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "收藏成功",
	})

	return
}

func UncollectContent(c *gin.Context) {
	var (
		usercollectcontent models.UserCollectContent
	)

	if err := c.ShouldBindJSON(&usercollectcontent); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("err:", err)
		return
	}

	err := models.UncollectContent(usercollectcontent.UserID, usercollectcontent.ContentID)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code":    0,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "取消收藏成功",
	})

	return
}
