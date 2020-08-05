package processexe

import (
    // "time"
	// "fmt"
    "log"
)

// type ServiceOutput struct {
// 	Outputname string
// 	Outputdata string
// }
// map[string]string{"Outputname": "Outputdata", "Outputname": "Outputdata"}


func Executeprocess(data4output map[string]string) int  {
	log.Println(">>> start to execute process <<<")
	processid_arr :=  processMap_ptr.getAllkeys()
	for _, processid_i := range processid_arr {
		ing_Process, _ := processMap_ptr.getProcessEleByID(processid_i)
		log.Printf(">>> deal with process %s ", processid_i)
		if ( ing_Process.Status == "running" ) {
			nexthop_arr := ing_Process.Nexthop
			for _, nexthop_i := range nexthop_arr {
				// 拼接发给下一跳的参数
				url_string := nexthop_i.NextService + "?info=" + ing_Process.Info
				for outport_i, nextport_i := range (nexthop_i.OutPort_NextPort) {
					log.Printf("nextport_i = %s, outport_i = %s", nextport_i, outport_i)
					url_string = url_string + "&" + nextport_i + "=" + data4output[outport_i]
				}

				// 发送给下一跳
				log.Println(url_string)

			}
		}

	}
	return 0
}