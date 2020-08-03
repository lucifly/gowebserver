package processexe

import (
    "time"
	// "fmt"
    "log"
)

// 开始执行业务流程
/*
* func StartProcess(processid int, intav int, maxcount int) int 
* @param 业务流程ID, 执行时间间隔, 总执行次数
* @return 执行结果，状态码
*/
func StartProcess(processid int, intav int, maxcount int) int {
	ch := make(chan int, 1)
	executionservice(intav , maxcount , ch ) 
}

func executionservice(intav int, maxcount int, ch chan int) int {
	

	timeout := time.NewTimer(time.Microsecond * 500)

	select {
	case x = <-ch:
		// 接收，信号
		return x, nil
	case <-timeout.C:
		//执行具体的业务流程

		/*
		   可以另外引入 service 执行的包
		*/

		return 0, errors.New("read time out")
	}

}