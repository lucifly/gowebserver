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
    NextPort string // 下一跳服务接收的接口
	OutPort string //当前服务的输出接口
}

// 声明processMap
var processMap map[string]ProcessTableEle


func getOneProcessTableEle() ProcessTableEle  {
	nexthop_1 := NextHop{ NextService:"2", NextPort:"n1", OutPort:"o1" }
	nethop_array := [] NextHop { nexthop_1, } 
	// bufCh := make(chan int, 1)
	processEle := ProcessTableEle{ ProcessID:"p1", Nexthop:nethop_array, Info:"info1", Csign:nil } 
	return processEle
}


func init()  {
	log.Println(">>>> process executor init <<<<")
	/* 使用 make 函数 */
	processMap = make(map[string]ProcessTableEle)

	var ing_Process ProcessTableEle
	ing_Process = getOneProcessTableEle()

	processMap[ing_Process.ProcessID] = ing_Process
}

// 开始执行业务流程
/*
* func StartProcess(processid int, intav int, maxcount int) int 
* @param 业务流程ID, 执行时间间隔, 总执行次数
* @return 执行结果，状态码
*/
func StartProcess(processid string, intav int, maxcount int) int {
    /*查看元素在集合中是否存在 */
    ing_Process, ok := processMap [ processid ] /*如果确定是真实的,则存在,否则不存在 */
    if (ok) {
		log.Printf("ready to execute %s ", processid)
		ch := make(chan int, 1)
		ing_Process.Csign = ch
		ing_Process.Status = "running"
		processMap[ing_Process.ProcessID] = ing_Process 
		// 注意，map元素是无法取址的，不可以直接用map["key"].value = newvalue来直接修改
		// 需要再写回map
		go executionservice(intav , maxcount , ing_Process ) 
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
    ing_Process, ok := processMap [ processid ] /*如果确定是真实的,则存在,否则不存在 */
    if (ok) {
		log.Printf("status of %s is %s", processid, ing_Process.Status)
		log.Printf("stop %s ", processid)
		ing_Process.Status = "stop"
		processMap[ing_Process.ProcessID] = ing_Process
		ing_Process.Csign <- 5
		return 0
    } else {
		log.Printf("process %s not exist", processid)
		return -1
    }
}

func executionservice(intav int, maxcount int, process_now ProcessTableEle) int {
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
	return 0	
}