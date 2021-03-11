package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"shuwo/models"
)

func UserLogin(c *gin.Context) {
	var (
		res = gin.H{}
	)

	user := &models.User{}
	userdb := &models.User{}

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
	}

	if user.UserName == "" || user.UserOpenid == "" {
		fmt.Println("用户名或密码不能为空!")
		res["message"] = "用户名或密码不能为空!"
		return
	}

	userdb, err = models.UserLogin(user.UserName, user.UserOpenid)

	if gorm.IsRecordNotFoundError(err) {
		fmt.Println("查询返回结果为空")
	}

	if err != nil || user.UserOpenid != userdb.UserOpenid || user.UserName != userdb.UserName {
		fmt.Println("用户名密码错误!")
		return
	}

	fmt.Println("用户名密码正确")

	return

}

func AppletsUserInfo(c *gin.Context) {
	var (
		user = &models.User{}
		res  = gin.H{}
	)

	if err := c.ShouldBindJSON(&user); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.UserName == "" || user.UserOpenid == "" {
		fmt.Println("用户名或密码不能为空!")
		res["message"] = "用户名或密码不能为空!"
		return
	}

	user, err := models.AppletsUserInfo(user.UserOpenid, user.UserName)

	if err != nil {
		res["message"] = err.Error()
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"user": map[string]interface{}{
			"UserOpenid": user.UserOpenid,
			"UserName":   user.UserName,
		},
	})

}

func LoginApplets(c *gin.Context) {
	var (
		user = models.User{}
		res  = gin.H{}
	)

	if err := c.ShouldBindJSON(&user); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.GetUser(user.UserOpenid)
	if err != nil {
		res["message"] = err.Error()
		res["code"] = 0
		log.Println(err)
		return
	}
	res["code"] = 1
}
