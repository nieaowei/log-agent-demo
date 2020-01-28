/*******************************************************
 *  File        :   tailf.go
 *  Author      :   nieaowei
 *  Date        :   2020/1/25 10:44 下午
 *  Notes       :
 *******************************************************/
package tailf

import (
	"errors"
	"github.com/gogf/gf/frame/g"
	"github.com/hpcloud/tail"
	"log-agent-demo/instance"
)

//type InstManager struct {
//	Insts []*TailfInst
//}

type TailfInst struct {
	Name string
	*tail.Tail
	ch chan *instance.Message
}

//var InstMgr *InstManager = &InstManager{}

//func init() {
//	path := g.Cfg().GetString("tailf.path")
//	suffix := g.Cfg().GetString("tailf.suffix")
//
//	files, err := ioutil.ReadDir(path)
//	if err != nil {
//		g.Log().Fatal("scan log file failed.")
//	}
//
//	for _, file := range files {
//		if file.IsDir() {
//			files1, err := ioutil.ReadDir(filepath.Join(path, "/", file.Name()))
//			if err != nil {
//				continue
//			}
//			for _, file1 := range files1 {
//				if strings.HasSuffix(file1.Name(), suffix) {
//					tailfInst, err := NewTailfInst(path+"/"+file1.Name(), DefaultConfig())
//					if err != nil {
//						g.Log().Warning("new inst faile.")
//						continue
//					}
//					tailfInst.Name = file.Name()
//					InstMgr.Insts = append(InstMgr.Insts, &tailfInst)
//				}
//			}
//		}
//	}
//}

func DefaultConfig() (config tail.Config) {
	return tail.Config{
		Location: &tail.SeekInfo{
			Offset: 0,
			Whence: 2,
		},
		ReOpen:      true,
		MustExist:   true,
		Poll:        true,
		Pipe:        false,
		RateLimiter: nil,
		Follow:      true,
		MaxLineSize: 0,
		Logger:      nil,
	}
}

func NewTailfInst(fileName string, config tail.Config) (inst *TailfInst, err error) {
	inst = &TailfInst{}
	inst.Tail, err = tail.TailFile(fileName, config)
	if err != nil {
		return
	}
	inst.Name = "default"
	inst.ch = make(chan *instance.Message, 10)
	return
}

//func (p *InstManager) LoadConfig() {
//	for _, inst := range p.Insts {
//		inst.LoadConfig()
//	}
//}

func NewInstance(fileName string, config tail.Config) (inst instance.Instance, err error) {
	inst1 := &TailfInst{}
	inst1.Tail, err = tail.TailFile(fileName, config)
	if err != nil {
		return
	}
	inst1.Name = "default"
	inst1.ch = make(chan *instance.Message, 10)
	return inst1, nil
}

func (p *TailfInst) LoadConfig() {
	p.ReOpen = g.Cfg().GetBool("tailf.ReOpen")
	p.MustExist = g.Cfg().GetBool("tailf.MustExist")
	p.Poll = g.Cfg().GetBool("tailf.Poll")
	p.Pipe = g.Cfg().GetBool("tailf.Pipe")
	p.Follow = g.Cfg().GetBool("tailf.Follow")
	p.MaxLineSize = g.Cfg().GetInt("tailf.MaxLineSize")
}

func (p *TailfInst) SendMsg(msg *instance.Message) {
	g.Log().Notice("tailf send msg", msg)
	p.ch <- msg
	return
}

func (p *TailfInst) ReceMsg() (msg *instance.Message, err error) {
	line, ok := <-p.Lines
	if !ok {
		g.Log().Notice(p.Filename + "channel closed.")
		return nil, errors.New("channel closed")
	}
	msg = &instance.Message{
		Topic: p.Name,
		Text:  line.Text,
	}
	g.Log().Notice("tailf rece msg", msg)
	return
}

func (p *TailfInst) Exce() {
	for {
		msg, err := p.ReceMsg()
		if err != nil {
			g.Log().Warning(err)
			break
		}
		p.SendMsg(msg)
	}
}
