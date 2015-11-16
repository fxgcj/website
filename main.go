package main

import (
	"github.com/astaxie/beego"
	"github.com/fxgcj/website/conf"
	logpkg "github.com/fxgcj/website/lib/log"
	"github.com/fxgcj/website/routers"
)

var log = logpkg.GetLogger()

func init() {
	conf.LoadConf("conf/v2.json")
}

func main() {
	BeegoInit()
	beego.Run()
}

func BeegoInit() {
	beego.TemplateLeft = "<<<"
	beego.TemplateRight = ">>>"
	routers.LoadRouters()
}
