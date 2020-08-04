# 2020-07-31
完成修改，从config.json中读取配置信息。  
Tips, 增加配置项需要在config.go中同步增加struct Config中的项。  

添加特性，添加命令行参数

# 2020-08-01
~~添加process处理的API接口~~  
    ~~`/process/deploy` // 新增 process~~  
    ~~`/process/delet` // 删除 process~~  
    ~~`/process/start` // 开始 process~~  
    ~~`/process/stop` // 结束 process~~  
    ~~`/process/status` // 获取各个process的状态~~ 

# 2020-08-03
设计了process执行的流程，以及信号的接入  
写了一个新的包 `processexe` 来负责process的五种接口的逻辑  

# 2020-08-04
## 修改process处理的API接口为   
    `/processcontrol/deploy` // 新增 process  
    `/processcontrol/delet` // 删除 process  
    `/processcontrol/start` // 开始 process  
    `/processcontrol/stop` // 结束 process  
    `/processcontrol/status` // 获取各个process的状态 

- `/processcontrol/start?processid=p1&intav=10&maxcount=4` 
开始执行某 process;  
processid=process ID, intav=间隔时间，maxcount=执行次数


- `/processcontrol/stop?processid=p1` 
提前结束某 process;  
processid=process ID

## 添加process记录table
```go
type ProcessTableEle struct { // 记录每一项process
	ProcessID string // process ID
	Nexthop []NextHop // 输出发送给下一跳
    Info string // process当前service所需的参数
    Status string // 当前service在process的状态
	Csign chan int // 用于对当前service发送 信号 的channel
}

type NextHop struct { // 下一跳记录结构
    NextService string // 下一跳的服务名
    NextPort string // 下一跳服务接收的接口
	OutPort string //当前服务的输出接口
}

processMap = make(map[string]ProcessTableEle)

// 注意，map元素是无法取址的，不可以直接用map["key"].value = newvalue来直接修改
// 需要再写回map. map["key"] = temp; temp.value = newvalue; map["key"] = temp
```