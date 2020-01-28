/*******************************************************
 *  File        :   config.go
 *  Author      :   nieaowei
 *  Date        :   2020/1/25 10:14 下午
 *  Notes       :	全局系统配置管理。
 *******************************************************/
package config

import (
	"github.com/gogf/gf/frame/g"
)

//type Configed interface {
//	LoadConfig()
//}
//
//// ReloadConfig is to reload the configuation of instance that
//// has implemented the interface.
//func ReloadConfig(inst interface{}) {
//	configed := inst.(Configed)
//	configed.LoadConfig()
//	g.Log().Debug("reload config success.")
//}

// ConfigFileInit is to set config file.
func configFileInit() {
	g.Cfg().SetFileName("config/config.ini")
	g.Log().Debug("Log and config manager init successed.")
}

func init() {
	configFileInit()
}
