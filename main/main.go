/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2020/1/25 10:10 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
	_ "log-agent-demo/config"
	"log-agent-demo/logs/tailf"
)

func main() {
	var waitSign chan int
	fmt.Println(tailf.InstMgr.Insts)
	for _, inst := range tailf.InstMgr.Insts {
		inst.Run()
	}
	for v := range waitSign {
		if v == 1 {
			break
		}
	}
}
