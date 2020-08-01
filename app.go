// main.go
package main

import (
	"io"
	"net/http"
	"os"
	"log"
	"flag"
	"strconv"
	"webserverRaspberry/config"
	//使用go mode 路径必须用当前项目名开头
	"webserverRaspberry/routers"

	"github.com/gin-gonic/gin"
)

func usage() {
	log.Println("version: 1.0.0")
	log.Println("Usage: app [-hpb]")
	log.Println("Options:")
	flag.PrintDefaults()
}

func main() {


	JsonParse := config.NewJsonStruct()
    v := config.Config{}
    //下面使用的是相对路径，config.json文件和main.go文件处于同一目录下
	JsonParse.Load("./config/config.json", &v)
	
    // 定义几个变量，用于接收命令行的参数值
    var webListenPort string
	var heartBeat    int
	var h bool
	// &user 就是接收命令行中输入 -u 后面的参数值，其他同理
	
	flag.BoolVar(&h, "h", false, "help")
    flag.StringVar(&webListenPort, "p", v.ListenPort, "web监听的端口，默认为:"+ v.ListenPort)
    flag.IntVar(&heartBeat, "b", 0, "心跳包频率 /ms，默认为空，不发送心跳包")
    // 解析命令行参数写入注册的flag里
    flag.Parse()

	flag.Usage = usage

	if h {
		flag.Usage()
		return
	}

	log.Println( "web listen at: " + webListenPort)
	log.Println( "heartBeat every [" + strconv.Itoa(heartBeat) + "] ms")

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
	process := r.Group("/process")
	{
		process.GET("/deploy", routers.DeployProcess) // 新增 process
		process.GET("/delet", routers.DeletProcess) // 删除 process
		process.GET("/start", routers.StartProcess) // 开始 process
		process.GET("/stop", routers.StopProcess) // 结束 process
		process.GET("/status", routers.StatusProcess) // 获取各个process的状态
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
