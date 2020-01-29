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
	LoadConfig()
	SendMsg(msg *Message)
	ReceMsg() (msg *Message, err error)
	BindChan(msgCh chan *Message)
}

type InstanceManager interface {
	RunAll()
	Register(inst Instance)
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

// default Instance manager.
var InstMgr *instMgr = &instMgr{}

func (p *instMgr) Register(inst Instance) {
	p.insts = append(p.insts, inst)
	p.Number++
	g.Log().Debug("Registered instance :", p.Number)
}

func (p *instMgr) RunAll() {
	for _, inst := range p.insts {
		go p.Exce(inst)
	}
}

func (p *instMgr) Exce(inst Instance) {
	for {
		msg, err := inst.ReceMsg()
		if err != nil {
			g.Log().Warning(err)
			break
		}
		inst.SendMsg(msg)
	}
}

func init() {
	//InstMgr = &instMgr{}
}
