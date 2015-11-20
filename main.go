package main

import (
	"github.com/astaxie/beego"
	"github.com/fxgcj/website/conf"
	logpkg "github.com/fxgcj/website/lib/log"
	"github.com/fxgcj/website/routers"
)

var log = logpkg.GetLogger()
var config *conf.Config

func init() {
	var err error
	config, err = conf.LoadConf("conf/v2.json")
	if err != nil {
		panic(err)
	}
	BeegoInit()
}

func main() {
	beego.Run()
}

func BeegoInit() {
	beego.RunMode = config.AppConfig.RunMode
	beego.AppName = config.AppConfig.Name
	beego.HttpPort = config.AppConfig.Port
	beego.BeegoServerName = config.AppConfig.ServerName
	beego.SessionOn = true
	beego.SessionName = "ckeyer"
	// beego.SessionDomain = "fxgcj.org"
	beego.SessionAutoSetCookie = true

	routers.LoadRouters()
}
