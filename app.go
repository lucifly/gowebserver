// main.go
package main

import (
	"io"
	"net/http"
	"os"
	"log"
	"webserverRaspberry/config"
	//使用go mode 路径必须用当前项目名开头
	"webserverRaspberry/routers"

	"github.com/gin-gonic/gin"
)

func main() {


	JsonParse := config.NewJsonStruct()
    v := config.Config{}
    //下面使用的是相对路径，config.json文件和main.go文件处于同一目录下
	JsonParse.Load("./config/config.json", &v)
	
	webListenPort := v.ListenPort

    log.Println( "web listen at: " + webListenPort)

	// 创建记录日志的文件
	f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	// 第一个参数是api 第二个静态文件的文件夹相对目录
	r.StaticFS("/static", http.Dir("./static"))

	// 第一个参数是api 第二个参数是具体的文件名字
	// r.StaticFile("/favicon.ico", "./static/img/favicon.ico")

	r.GET("/", hello)

	// 路由组
	data := r.Group("/data")
	{
		data.GET("/rasdata", routers.RasData)
		data.GET("/webdata", routers.WebData)
		//可以自己添加其他，一个请求的路径对应一个函数

		// ...
	}

	r.Run(":"+webListenPort) // listen and serve on 0.0.0.0:8080
}

func hello(context *gin.Context) {
	println(">>>> hello function start <<<<")
	context.Redirect(http.StatusMovedPermanently, "http://47.98.169.4:6082/static/dashboard/index.html")
	// context.JSON(http.StatusOK, gin.H{
	// 	"code":    200,
	// 	"success": true,
	// })
}
