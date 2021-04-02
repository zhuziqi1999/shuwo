package main

import (
	"fmt"
	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"github.com/zhuziqi1999/shuwo/controllers"
	"github.com/zhuziqi1999/shuwo/models"
	"github.com/zhuziqi1999/shuwo/system"

	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//configFilePath := flag.String("C", "conf/conf.yaml", "config file path")
	//logConfigPath := flag.String("L", "conf/seelog.xml", "log config file path")
	//flag.Parse()
	//
	///*log*/
	//logger, err := seelog.LoggerFromConfigAsFile(*logConfigPath)
	//if err != nil {
	//	seelog.Critical("err parsing seelog config file", err)
	//	return
	//}
	//seelog.ReplaceLogger(logger)
	//defer seelog.Flush()
	//
	//if err := system.LoadConfiguration(*configFilePath); err != nil {
	//	seelog.Critical("err parsing config log file", err)
	//	return
	//}

	/*database*/
	db, err := models.InitDB()
	if err != nil {
		seelog.Critical("err open databases", err)
		return
	}
	fmt.Println("dbs success")
	defer db.Close()

	/*yaml*/
	cluster := new(system.Configuration)
	yamlFile, err := ioutil.ReadFile("conf/conf.yaml")
	if err != nil {
		log.Println(err)
	}

	err = yaml.Unmarshal(yamlFile, &cluster)

	if err != nil {
		log.Println(err)
	}

	/*gin*/
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.POST("/hello", func(c *gin.Context) {
		fmt.Println("hello!")
	})

	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	//user
	router.POST("/appletsUserInfo", controllers.AppletsUserInfo)
	router.POST("/loginApplets", controllers.LoginApplets)

	//content
	router.POST("/createContent", controllers.CreateContent)
	router.POST("/getContentList", controllers.GetContentList)
	router.POST("/likeContent", controllers.LikeContent)
	router.POST("/unlikeContent", controllers.UnlikeContent)
	router.POST("/getLikeList", controllers.GetLikeList)
	router.POST("/collectContent", controllers.CollectContent)
	router.POST("/uncollectContent", controllers.UncollectContent)

	//group
	router.POST("/createGroup", controllers.CreateGroup)
	router.POST("/getGroupList", controllers.GetGroupList)
	router.POST("/inGroup", controllers.InGroup)
	router.POST("/outGroup", controllers.OutGroup)
	router.POST("/getMyGroupList", controllers.GetMyGroupList)

	//comment
	router.POST("/createComment", controllers.CreateComment)
	router.POST("/getCommentList", controllers.GetCommentList)

	//file
	router.POST("/upload", controllers.UploadFile)

	//folder
	router.POST("/createFolder", controllers.CreateFolder)
	router.POST("/getFolderList", controllers.GetFolderList)
	router.POST("/deleteFolder", controllers.DeleteFolder)

	//message
	router.POST("/getLikeMessageList", controllers.GetLikeMessageList)
	router.POST("/getCommentMessageList", controllers.GetCommentMessageList)
	router.POST("/getCollectMessageList", controllers.GetCollectMessageList)

	http.ListenAndServe(cluster.Addr, router)

}
