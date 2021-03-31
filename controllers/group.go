package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhuziqi1999/shuwo/models"
	"log"
	"net/http"
)

func CreateGroup(c *gin.Context) {
	var (
		group = &models.Group{}
		res   = gin.H{}
	)

	if err := c.ShouldBindJSON(&group); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("err:", err)
		return
	}

	group, err := models.CreateGroup(group.GroupCreatedBy, group.GroupName, group.GroupRemark)

	if err != nil {
		log.Println(err)
		res["message"] = err.Error()
		return
	}

	err = models.InGroup(group.GroupCreatedBy, group.GroupID)

	if err != nil {
		log.Println(err)
		res["message"] = err.Error()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"user": map[string]interface{}{
			"GroupID":   group.GroupID,
			"GroupName": group.GroupName,
		},
	})

}

func GetGroupList(c *gin.Context) {
	var (
		group interface{}
		user  = &models.User{}
	)

	if err := c.ShouldBindJSON(&user); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("err:", err)
		return
	}

	group = models.GetGroupList(user.UserOpenid)

	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"content": group,
	})

}

func InGroup(c *gin.Context) {
	var (
		usergroup models.UserGroup
	)

	if err := c.ShouldBindJSON(&usergroup); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("err:", err)
		return
	}

	err := models.InGroup(usergroup.UserID, usergroup.GroupID)

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
		"message": "加入成功",
	})

	return
}

func OutGroup(c *gin.Context) {
	var (
		usergroup models.UserGroup
	)

	if err := c.ShouldBindJSON(&usergroup); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("err:", err)
		return
	}

	err := models.OutGroup(usergroup.UserID, usergroup.GroupID)

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
		"message": "退出成功",
	})

	return
}

func GetMyGroupList(c *gin.Context) {
	var (
		group interface{}
		user  = &models.User{}
	)

	if err := c.ShouldBindJSON(&user); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("err:", err)
		return
	}

	group = models.GetMyGroupList(user.UserOpenid)

	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"content": group,
	})

}
