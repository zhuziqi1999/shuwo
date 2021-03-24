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

	router.POST("/appletsUserInfo", controllers.AppletsUserInfo)
	router.POST("/loginApplets", controllers.LoginApplets)
	router.POST("/createContent", controllers.CreateContent)
	router.POST("/getHotContentList", controllers.GetHotContentList)
	router.POST("/createGroup", controllers.CreateGroup)
	router.POST("/getGroupList", controllers.GetGroupList)

	http.ListenAndServe(cluster.Addr, router)

}
