package controllers

import (
	_ "container/list"

	"github.com/fxgcj/website/conf"
)

type BaseController struct {
	Controller
}

func (b *BaseController) Prepare() {
	b.InitWebPage()
}

func (b *BaseController) InitWebPage() {
	website = conf.GetConf().WebSite

	b.Ctx.Request.Header.Add("Access-Control-Allow-Origin", "*")

	// 验证是否来自合法域名访问
	if !b.IsAllowHost() {
		b.Ctx.WriteString(`<!DOCTYPE html><html><head><meta http-equiv="refresh" content="0; url=` +
			website.HostUrl + string([]byte(b.Ctx.Input.Url())) + `" /></head></html>`)
		b.StopRun()
	}

	b.Data["WebSite"] = website
	b.Data["WebsiteName"] = website.Title
	b.Data["HostUrl"] = website.HostUrl
	b.SetPageTitle("首页")
	b.SetDescript(website.Description)
	b.AddKeyWord(website.Keywords...)

	b.Data["Metes"] = ""
	b.AddCustomCssStyle("//cdn.bootcss.com/bootstrap/3.3.5/css/", "bootstrap.min.css", "bootstrap-theme.min.css")
	b.AddCustomCssStyle("//cdn.bootcss.com/font-awesome/4.4.0/css/", "font-awesome.min.css")
	b.AddCssStyle("style.css")
	//	b.AddCustomCssStyle("http://fonts.useso.com/", "css?family=Open+Sans:300,400,600&subset=latin,latin-ext")

	b.AddCustomJsScript("//cdn.bootcss.com/jquery/2.1.4/", "jquery.min.js")
	b.AddCustomJsScript("//cdn.bootcss.com/bootstrap/3.3.5/js/", "bootstrap.min.js")
	b.AddCustomJsScript("//cdn.bootcss.com/jquery-migrate/1.2.1/", "jquery-migrate.min.js")
	b.AddCustomJsScript("//cdn.bootcss.com/wow/1.1.2/", "wow.min.js")

	b.Data["Tail"] = `热眼看社会，冷眼看风险。`

	b.LayoutSections = make(map[string]string)

	b.Layout = "layout/index.html"
}
