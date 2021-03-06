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
写了一个新的包 `processexe` 来负责process的~~五~~四种接口的逻辑  

# 2020-08-04
## 修改process处理的API接口为 **processcontrol** 接口
   ~~`/processcontrol/deploy` // 新增 process~~  
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

- `/processcontrol/delet?processid=p1` 
删除某 process;  
processid=process ID

## ~~添加process记录table~~


# 2020-08-05
删除接口 ~~`/processcontrol/deploy`~~ 。  
业务逻辑改为，  
1. 在接收到 `/processcontrol/start` 命令的时候，会查询本地processTable，如果没有，查询远程业务流程数据库中的，如果也没有就报错。  
2. 而新增的业务流程只存在远程业务流程数据库中。本地要用的时候才会去查询远程业务流程数据库。
  

完成剩余的 `/processcontrol/start`，`/processcontrol/stop`，`/processcontrol/delet` 接口

完成processexecution  
1. 在 `\data` 下添加 `\data\processdata` 接口，用于接收process的数据
## 修改 processTable 的数据结构
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
	OutPort_NextPort map[string]string // 输出与发给下一跳的接口的键值对
	// "outport":"nextport"
}

// 命名ProcessMap类型，并添加相应的方法
type ProcessMap map[string]ProcessTableEle

// 声明processMap指针
var processMap_ptr *ProcessMap

// 生成缺省的processMap 用于调试
func NewDefaultProcessMap() *ProcessMap  {

// 根据process id获取 当前process
func (this *ProcessMap)  getProcessEleByID(processid string) (ProcessTableEle, bool)  {

// 编辑指定process id 的process
func (this *ProcessMap)  editProcessEleByID(processid string, new_process ProcessTableEle) (bool)  {

// 根据process id 删除process
func (this *ProcessMap)  delProcessEleByID(processid string) (bool)  {

// 获取所有的process id
func (this *ProcessMap) getAllkeys() []string  {

```