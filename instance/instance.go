/*******************************************************
 *  File        :   inst.go
 *  Author      :   nieaowei
 *  Date        :   2020/1/27 2:23 上午
 *  Notes       :	Instance manager.
					1. Manages registered instances.
					2.
 *******************************************************/
package instance

import "github.com/gogf/gf/frame/g"

// Other instances that have implemented the interface
// can be uniformly scheduled by the instance manager.
type Instance interface {
	Exce()
	LoadConfig()
	SendMsg(msg *Message)
	ReceMsg() (msg *Message, err error)
}

// Common communication protocol.
type Message struct {
	Topic string
	Text  string
}

// default Instance manager.
type instMgr struct {
	insts  []Instance
	Number int
}

type InstanceManager interface {
	RunAll()
	Register(inst Instance)
}

var InstMgr *instMgr = &instMgr{}

func (p *instMgr) Register(inst Instance) {
	p.insts = append(p.insts, inst)
	p.Number++
	g.Log().Debug("Registered instance :", p.Number)
}

func (p *instMgr) RunAll() {
	for _, inst := range p.insts {
		go inst.Exce()
	}
}

func init() {
	//InstMgr = &instMgr{}
}
