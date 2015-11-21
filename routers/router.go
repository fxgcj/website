package routers

import (
	"github.com/astaxie/beego"
	"github.com/fxgcj/website/controllers"
	logpkg "github.com/fxgcj/website/lib/log"
)

var log = logpkg.GetLogger()

func LoadRouters() {
	log.Info("加载路由信息")

	// beego.SetStaticPath("/img", "blog/img")
	beego.AutoRouter(&controllers.AdminController{})
	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/blog/:id", &controllers.BlogController{})
	beego.Router("/tag", &controllers.TagController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/:year:int/:month:int", &controllers.ArchiveController{})

}
