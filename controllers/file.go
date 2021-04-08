package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhuziqi1999/shuwo/models"
	"net/http"
	"path/filepath"
	"strings"
)

func UploadFile(c *gin.Context) {
	basePath := "C:/gin-upload/"
	name := c.PostForm("filename")
	openid := c.PostForm("filecreatedby")
	folderid := c.PostForm("filefolderid")

	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	//获取后缀
	names := strings.Split(file.Filename, ".")
	filetype := names[len(names)-1]
	fmt.Println(filetype)

	fileid, err := models.CreateFile(openid, name, file.Size, filetype, basePath, folderid)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code":    0,
			"message": err,
		})
		return
	}

	fname := fileid + "." + filetype
	filename := basePath + filepath.Base(fname)

	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code":    0,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "上传成功",
	})
}

func GetFileList(c *gin.Context) {
	var (
		file     = &models.File{}
		filelist interface{}
	)

	if err := c.ShouldBindJSON(&file); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("err:", err)
		return
	}

	filelist = models.GetFileList(file.FileCreatedBy, file.FileFolderID)

	c.JSON(http.StatusOK, gin.H{
		"code":     1,
		"filelist": filelist,
	})

}

func DeleteFile(c *gin.Context) {
	var (
		file models.File
	)

	if err := c.ShouldBindJSON(&file); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("err:", err)
		return
	}

	err := models.DeleteFile(file.FileCreatedBy, file.FileID)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
	})
	return

}
