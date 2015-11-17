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
	// beego.Router("/", &controllers.IndexController{})
	beego.Router("/:name:string.html", &controllers.BlogController{})
	// beego.Router("/tag", &controllers.TagController{})
	// beego.Router("/category", &controllers.CategoryController{})
	// beego.Router("/archive/:year:string-:month:string.html", &controllers.ArchiveController{})

	// beego.Router("/webhook", &controllers.WebhookController{})
}
