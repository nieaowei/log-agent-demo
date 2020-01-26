/*******************************************************
 *  File        :   tailf.go
 *  Author      :   nieaowei
 *  Date        :   2020/1/25 10:44 下午
 *  Notes       :
 *******************************************************/
package tailf

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/hpcloud/tail"
	"os"
	"strings"
)

type InstManager struct {
	Insts []*TailfInst
}

type TailfInst struct {
	Name string
	*tail.Tail
}

var InstMgr *InstManager = &InstManager{}

func init() {
	path := g.Cfg().GetString("tailf.path")
	suffix := g.Cfg().GetString("tailf.suffix")
	files, err := os.Open(path)
	if err != nil || files == nil {
		g.Log().Fatal("scan log file failed.")
	}
	filesNames, err := files.Readdirnames(0)
	if err != nil {
		g.Log().Fatal("scan log file failed.")
	}
	for _, fileName := range filesNames {
		if strings.HasSuffix(fileName, suffix) {
			tailfInst, err := NewTailfInst(path+"/"+fileName, defaultConfig())
			if err != nil {
				g.Log().Warning("new inst faile.")
				continue
			}
			InstMgr.Insts = append(InstMgr.Insts, &tailfInst)
		}
	}
}

func defaultConfig() (config tail.Config) {
	return tail.Config{
		Location:    nil,
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

func NewTailfInst(fileName string, config tail.Config) (inst TailfInst, err error) {
	inst.Tail, err = tail.TailFile(fileName, config)
	if err != nil {
		return
	}
	return
}

func (p *InstManager) LoadConfig() {
	for _, inst := range p.Insts {
		inst.LoadConfig()
	}
}

func (p *TailfInst) LoadConfig() {
	p.ReOpen = g.Cfg().GetBool("tailf.ReOpen")
	p.MustExist = g.Cfg().GetBool("tailf.MustExist")
	p.Poll = g.Cfg().GetBool("tailf.Poll")
	p.Pipe = g.Cfg().GetBool("tailf.Pipe")
	p.Follow = g.Cfg().GetBool("tailf.Follow")
	p.MaxLineSize = g.Cfg().GetInt("tailf.MaxLineSize")
}

func (p *TailfInst) Run() {
	go func() {
		for {
			line := <-p.Lines
			fmt.Println(line.Text)
		}
	}()
}
