package config

import (
    "io/ioutil"
    "encoding/json"
	// "fmt"
    "log"
)
//定义配置文件解析后的结构
type Config struct {
    ListenPort  string `json:"listenPort"`
    WsPort string `json:"wsPort,omitempty"`
    Servicename string `json:"servicename,omitempty"`
    DataAPI string `json:"dataAPI,omitempty"`
}
// 给结构体的每个字段打上一个标签，标签冒号前是类型，后面是标签名
// tag里面加上omitempy，可以在序列化的时候忽略0值或者空值
// 结构体类型和需要的类型不一致，还可以指定,支持string,number和boolean
// Number    int     `json:"number,string"`
// 参考博客 https://www.cnblogs.com/yorkyang/p/8990570.html 

func init() {
	log.Println(">>>> init config <<<<")
	// log.Println(">>>> get database connection start <<<<")
	// db = database.GetDataBase()
}

// func get() {
//     JsonParse := NewJsonStruct()
//     v := Config{}
//     //下面使用的是相对路径，config.json文件和main.go文件处于同一目录下
//     JsonParse.Load("./config.json", &v)
//     log.Println(v.ListenPort)
//     log.Println(v.WsPort)
// }

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
    return &JsonStruct{}
}

func (jst *JsonStruct) Load(filename string, v interface{}) {
    //ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return
    }

    //读取的数据为json格式，需要进行解码
    err = json.Unmarshal(data, v)
    if err != nil {
        return
    }
}