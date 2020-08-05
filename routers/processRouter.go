package routers

import (
	// "fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	// "program/com.ypc/helloGin/model"

	"webserverRaspberry/processexe"

)

// // RaspData 记录从树莓派传过来的数据
// type RaspData struct {
// 	id    string
// 	time  string
// 	tem   string
// 	hum   string
// 	shake string
// }

// // DataCache 暂时缓存树莓派的数据
// var DataCache RaspData

func init() {
	log.Println(">>>> process controller init <<<<")
	// log.Println(">>>> get database connection start <<<<")
	// db = database.GetDataBase()
}

// process := r.Group("/process")
// {
// 	data.GET("/deploy", routers.DeployProcess) // 新增 process
// 	data.GET("/delet", routers.DeletProcess) // 删除 process
// 	data.GET("/start", routers.StartProcess) // 开始 process
// 	data.GET("/stop", routers.StopProcess) // 结束 process
// 	data.GET("/status", routers.StatusProcess) // 获取各个process的状态


func DeployProcess(context *gin.Context) {
	// println(">>>> get DeployProcess <<<<")

	context.JSON(200, gin.H{
		"result": "ok",
	})

}

func DeletProcess(context *gin.Context) {
	// println(">>>> get DeletProcess <<<<")

	processid := context.Request.URL.Query().Get("processid")
	// log.Println( context.Request.URL.Query().Get("processid") )
	processexe.DelProcess(processid)

	context.JSON(200, gin.H{
		"result": "ok",
	})

}

// 开始执行某 process;  
// processid=process ID, intav=间隔时间，maxcount=执行次数
func StartProcess(context *gin.Context) {
	// println(">>>> get data from raspberry <<<<")
	// log.Println( context.Request.URL.Query().Get("tem") )
	processid := context.Request.URL.Query().Get("processid")
	intav, _ :=  strconv.Atoi(context.Request.URL.Query().Get("intav"))
	maxcount, _ :=  strconv.Atoi(context.Request.URL.Query().Get("maxcount"))

	processexe.StartProcess(processid, intav, maxcount)

	context.JSON(200, gin.H{
		"result": "ok",
	})

}

// 提前结束某 process;  
// processid=process ID
func StopProcess(context *gin.Context) {
	// println(">>>> get data from raspberry <<<<")
	processid := context.Request.URL.Query().Get("processid")
	// log.Println( context.Request.URL.Query().Get("tem") )
	processexe.StopProcess(processid)

	context.JSON(200, gin.H{
		"result": "ok",
	})

}

func StatusProcess(context *gin.Context) {
	// println(">>>> get data from raspberry <<<<")

	log.Println( context.Request.URL.Query().Get("tem") )

	context.JSON(200, gin.H{
		"result": "ok",
	})

}


// // RasData 接收从树莓派传过来的数据
// func RasData(context *gin.Context) {
// 	println(">>>> get data from raspberry <<<<")

// 	(DataCache).id = context.Query("id")
// 	(DataCache).tem = context.Request.URL.Query().Get("tem")
// 	(DataCache).hum = context.Request.URL.Query().Get("hum")
// 	(DataCache).shake = context.Request.URL.Query().Get("shake")
// 	(DataCache).time = time.Now().Format("2006-01-02 15:04:05")

// 	 log.Println("raspi get data from %s, data is %s\n", (DataCache).id, (DataCache).tem)

// 	context.JSON(200, gin.H{
// 		"result": "ok",
// 	})

// }

// // WebData 给前端返回数据
// func WebData(context *gin.Context) {
// 	println(">>>> Ruturn data to web <<<<")

// 	// (DataCache).id = context.Query("id")
// 	// (DataCache).tem = context.Request.URL.Query().Get("tem")
// 	// (DataCache).hum = context.Request.URL.Query().Get("hum")
// 	// (DataCache).shake = context.Request.URL.Query().Get("shake")
// 	// (DataCache).time = time.Now().Format("2006-01-02 15:04:05")

// 	 log.Println("web get data from %s, data is %s\n", (DataCache).id, (DataCache).tem)

// 	context.JSON(200, gin.H{
// 		"id":    (DataCache).id,
// 		"tem":   (DataCache).tem,
// 		"hum":   (DataCache).hum,
// 		"shake": (DataCache).shake,
// 		"time":  (DataCache).time,
// 	})

// }

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
