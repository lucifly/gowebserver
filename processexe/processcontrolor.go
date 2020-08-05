package processexe

import (
    "time"
	// "fmt"
    "log"
)

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
    // NextPort string // 下一跳服务接收的接口
	// OutPort string //当前服务的输出接口
}

type ProcessMap map[string]ProcessTableEle

// 声明processMap
var processMap_ptr *ProcessMap

// 生成缺省的processMap 用于调试
func NewDefaultProcessMap() *ProcessMap  {
	t_process := make(ProcessMap)

	var ing_Process ProcessTableEle
	ing_Process = getOneProcessTableEle()

	t_process[ing_Process.ProcessID] = ing_Process

	return &t_process

}

// 根据process id获取 当前process
func (this *ProcessMap)  getProcessEleByID(processid string) (ProcessTableEle, bool)  {
	tprocess,ok := (*this) [ processid ]
	if (ok) {
		return tprocess,ok
	} else {
		return tprocess,ok
	}
}

// 编辑指定process id 的process
func (this *ProcessMap)  editProcessEleByID(processid string, new_process ProcessTableEle) (bool)  {
	_,ok := (*this) [ processid ]
	if (ok) {
		(*this) [ processid ] = new_process
		return true
	} else {
		return false
	}
}

// 根据process id 删除process
func (this *ProcessMap)  delProcessEleByID(processid string) (bool)  {
	_,ok := (*this) [ processid ]
	if (ok) {
		delete(*this,processid)
		return true
	} else {
		return false
	}
}

// 获取所有的process id
func (this *ProcessMap) getAllkeys() []string  {

	keys := make([]string, len((*this)), len((*this)))
	k := 0
	for key, _ := range (*this) {
		keys[k] = key
		k++
	} 

	return keys
}

func getOneProcessTableEle() ProcessTableEle  {
	nexthop_1 := NextHop{ NextService:"2", OutPort_NextPort:map[string]string{"o1":"n1", "o2":"n2"} }
	nethop_array := [] NextHop { nexthop_1, } 
	// bufCh := make(chan int, 1)
	processEle := ProcessTableEle{ ProcessID:"p1", Nexthop:nethop_array, Info:"info1", Csign:nil } 
	return processEle
}


func init()  {
	log.Println(">>>> process executor init <<<<")
	/* 使用 make 函数 */
	processMap_ptr = NewDefaultProcessMap()

}

// 开始执行业务流程
/*
* func StartProcess(processid int, intav int, maxcount int) int 
* @param 业务流程ID, 执行时间间隔, 总执行次数
* @return 执行结果，状态码
*/
func StartProcess(processid string, intav int, maxcount int) int {
    /*查看元素在集合中是否存在 */
    ing_Process, ok := processMap_ptr.getProcessEleByID(processid) /*如果确定是真实的,则存在,否则不存在 */
    if (ok) {
		log.Printf("ready to execute %s ", processid)
		ch := make(chan int, 1)
		ing_Process.Csign = ch
		ing_Process.Status = "running"
		processMap_ptr.editProcessEleByID(ing_Process.ProcessID,ing_Process)
		// 注意，map元素是无法取址的，不可以直接用map["key"].value = newvalue来直接修改
		// 需要再写回map
		go startservice(intav , maxcount , ing_Process ) 
		return 0
    } else {
		log.Printf("process %s not exist", processid)
		return -1
    }
}

// 停止执行某业务流程
/*
* func StartProcess(processid int, intav int, maxcount int) int 
* @param 业务流程ID, 执行时间间隔, 总执行次数
* @return 执行结果，状态码
*/
func StopProcess(processid string) int {
    /*查看元素在集合中是否存在 */
	ing_Process, ok := processMap_ptr.getProcessEleByID(processid) /*如果确定是真实的,则存在,否则不存在 */
    if (ok) {
		log.Printf("status of %s is %s", processid, ing_Process.Status)
		if (ing_Process.Status != "running") {
			log.Printf("%s is not running", processid)
			return -2
		}
		log.Printf("stop %s ", processid)
		ing_Process.Status = "stop"
		processMap_ptr.editProcessEleByID(ing_Process.ProcessID,ing_Process)
		ing_Process.Csign <- 5
		return 0
    } else {
		log.Printf("process %s not exist", processid)
		return -1
    }
}


// 删除某业务流程
/*
* func StartProcess(processid int, intav int, maxcount int) int 
* @param 业务流程ID, 执行时间间隔, 总执行次数
* @return 执行结果，状态码
*/
func DelProcess(processid string) int {
    /*查看元素在集合中是否存在 */
	ing_Process, ok := processMap_ptr.getProcessEleByID(processid) /*如果确定是真实的,则存在,否则不存在 */
    if (ok) {
		log.Printf("status of %s is %s", processid, ing_Process.Status)
		if (ing_Process.Status != "running") {
			log.Printf("delet %s ", processid)
			processMap_ptr.delProcessEleByID(ing_Process.ProcessID)
			return 0
		}
		log.Printf("stop %s ", processid)
		ing_Process.Status = "stop"
		ing_Process.Csign <- 5
		log.Printf("delete %s ", processid)
		processMap_ptr.delProcessEleByID(ing_Process.ProcessID)
		return 0
    } else {
		log.Printf("process %s not exist", processid)
		return -1
    }
}

func startservice(intav int, maxcount int, process_now ProcessTableEle) int {
	log.Printf( "process start, execute every %d s, for %d times", intav, maxcount )
	for i := 0; i < maxcount; i++ {
		timeout := time.NewTimer(time.Second * time.Duration(intav) )
		select {
			case x := <-(process_now.Csign):
				// 接收，信号
				log.Printf("get sign %d, and exit", x)
				return 1
			case <-timeout.C:
				//执行具体的业务流程
				/*
				可以另外引入 service 执行的包
				*/
				log.Printf("timeout, timecount %d\n", i)			
		}	
	}
	log.Println("timeout over timecount")
	process_now.Status = "stop"
	processMap_ptr.editProcessEleByID(process_now.ProcessID,process_now)
	return 0	
}
