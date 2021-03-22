package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"shuwo/models"
)

func AppletsUserInfo(c *gin.Context) {
	var (
		user = &models.User{}
		res  = gin.H{}
	)

	if err := c.ShouldBindJSON(&user); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("err:", err)
		return
	}

	if user.UserName == "" || user.UserOpenid == "" {
		fmt.Println("用户名或openid不能为空!")
		res["message"] = "用户名或openid不能为空!"
		return
	}

	user, err := models.AppletsUserInfo(user.UserOpenid, user.UserName, user.UserAvatarUrl)

	if err != nil {
		res["message"] = err.Error()
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"user": map[string]interface{}{
			"UserOpenid": user.UserOpenid,
			"UserName":   user.UserName,
		},
	})
	return
}

func LoginApplets(c *gin.Context) {
	var (
		user = models.User{}
		res  = gin.H{}
	)

	if err := c.ShouldBindJSON(&user); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "code": 0})
		return
	}

	err := models.GetUser(user.UserOpenid)
	if err != nil {
		res["message"] = err.Error()
		res["code"] = 0
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"user": map[string]interface{}{
			"UserOpenid": user.UserOpenid,
			"UserName":   user.UserName,
		},
	})
}
