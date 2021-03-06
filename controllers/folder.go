package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhuziqi1999/shuwo/models"
	"log"
	"net/http"
)

func CreateFolder(c *gin.Context) {
	var (
		folder = &models.Folder{}
		res    = gin.H{}
	)

	if err := c.ShouldBindJSON(&folder); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("err:", err)
		return
	}

	folder, err := models.CreateFolder(folder.FolderCreatedBy, folder.FolderName, folder.FolderParentID)

	if err != nil {
		log.Println(err)
		res["message"] = err.Error()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"user": map[string]interface{}{
			"folderid":   folder.FolderID,
			"foldername": folder.FolderName,
		},
	})

}

func GetFolderList(c *gin.Context) {
	var (
		folder     = &models.Folder{}
		folderlist interface{}
	)

	if err := c.ShouldBindJSON(&folder); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("err:", err)
		return
	}

	folderlist = models.GetFolderList(folder.FolderCreatedBy, folder.FolderParentID)

	c.JSON(http.StatusOK, gin.H{
		"code":       1,
		"folderlist": folderlist,
	})

}

func DeleteFolder(c *gin.Context) {
	var (
		folder models.Folder
	)

	if err := c.ShouldBindJSON(&folder); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("err:", err)
		return
	}

	err := models.DeleteFolder(folder.FolderCreatedBy, folder.FolderID)

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
		})
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
		})
	}

}
