package routers

import (
	"github.com/astaxie/beego"
	"github.com/fxgcj/website/controllers"
	logpkg "github.com/fxgcj/website/lib/log"

	"html/template"
	"net/http"
)

var log = logpkg.GetLogger()

func LoadRouters() {
	log.Info("加载路由信息")

	beego.Errorhandler("404", NotFound)
	beego.Errorhandler("500", ServerErr)

	// beego.SetStaticPath("/img", "blog/img")
	beego.AutoRouter(&controllers.AdminController{})
	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/blog/:id", &controllers.BlogController{})
	beego.Router("/tag", &controllers.TagController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/:year:int/:month:int", &controllers.ArchiveController{})
	beego.Router("/search", &controllers.SearchController{})

}

var errtpl = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <title>Ckeyer - Error</title>
        <link rel="stylesheet" href="https://dn-ckeyer-css.qbox.me/404.css">
    </head>
    <body style="overflow: hidden;">
        
    </body>
	<body style="overflow: hidden;">
		<div id="container">
    		<ul id="scene" style="transform: translate3d(0px, 0px, 0px); transform-style: preserve-3d; backface-visibility: hidden;">
    		    <li class="layer" data-depth="0.10" style="position: relative; display: block; left: 0px; top: 0px; transform: translate3d(11.61px, -5.944px, 0px); transform-style: preserve-3d; backface-visibility: hidden;"><div class="star diamond"></div></li>
        		<li class="layer" data-depth="0.30" style="position: absolute; display: block; left: 0px; top: 0px; transform: translate3d(34.83px, -17.832px, 0px); transform-style: preserve-3d; backface-visibility: hidden;"><div class="star dot"></div></li>
        		<li class="layer" data-depth="0.50" style="position: absolute; display: block; left: 0px; top: 0px; transform: translate3d(58.05px, -29.72px, 0px); transform-style: preserve-3d; backface-visibility: hidden;"><div class="star sparkle"></div></li>
        		<li class="layer" data-depth="0.05" style="position: absolute; display: block; left: 0px; top: 0px; transform: translate3d(5.805px, -2.972px, 0px); transform-style: preserve-3d; backface-visibility: hidden;"><div class="lighthouse"></div></li>
        		<li class="layer" data-depth="0.20" style="position: absolute; display: block; left: 0px; top: 0px; transform: translate3d(23.22px, -11.888px, 0px); transform-style: preserve-3d; backface-visibility: hidden;"><div class="wave dark-blue depth-20"></div></li>
        		<li class="layer" data-depth="0.40" style="position: absolute; display: block; left: 0px; top: 0px; transform: translate3d(46.44px, -23.776px, 0px); transform-style: preserve-3d; backface-visibility: hidden;"><div class="wave medium-blue depth-40"></div></li>
        		<li class="layer" data-depth="0.60" style="position: absolute; display: block; left: 0px; top: 0px; transform: translate3d(69.66px, -35.664px, 0px); transform-style: preserve-3d; backface-visibility: hidden;"><div class="wave light-blue depth-60"></div></li>
        		<li class="layer" data-depth="0.00" style="position: absolute; display: block; left: 0px; top: 0px; transform: translate3d(0px, 0px, 0px); transform-style: preserve-3d; backface-visibility: hidden;">
        		    <div class="error-message">
        		    	<div align="center" ><font color="white" size=5 >观察君出海去了呢|~&>-。<font></div><br>
            		    	<div class="button"><a href="/">回到首页</a></div>
        		    </div>
    		    </li>
    		</ul>
        </div>
</body></html>
`

// show 404 notfound error.
func NotFound(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("beegoerrortemp").Parse(errtpl)
	// template.New("sd").ParseFiles(...)
	data := make(map[string]interface{})
	data["Title"] = "Page Not Found"
	data["Url"] = r.URL
	//rw.WriteHeader(http.StatusNotFound)
	t.Execute(rw, data)
}

func ServerErr(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("beegoerrortemp").Parse(errtpl)
	// template.New("sd").ParseFiles(...)
	data := make(map[string]interface{})
	data["Title"] = "观察君出海去了呢"
	data["Url"] = r.URL
	//rw.WriteHeader(http.StatusNotFound)
	t.Execute(rw, data)
}
