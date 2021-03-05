package controllers

import (
	"shuwo/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)



func UserLogin(c *gin.Context){
	var (
		res = gin.H{}
	)

	user := &models.User{}
	userdb := &models.User{}

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
	}

	if (user.UserName == "" || user.UserOpenid == "" ){
		fmt.Println("用户名或密码不能为空!")
		res["message"] ="用户名或密码不能为空!"
		return
	}

	userdb, err = models.UserLogin(user.UserName,user.UserOpenid)

	if gorm.IsRecordNotFoundError(err) {
		fmt.Println("查询返回结果为空")
	}


	if( err != nil || user.UserOpenid != userdb.UserOpenid || user.UserName != userdb.UserName) {
		fmt.Println("用户名密码错误!")
		return
	}

	fmt.Println("用户名密码正确")

	return

}
