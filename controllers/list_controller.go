package controllers

import (
	_ "container/list"

	"github.com/fxgcj/website/conf"
)

type ListController struct {
	Controller
}

func (l *ListController) Prepare() {
	website = conf.GetConf().WebSite

	l.Ctx.Request.Header.Add("Access-Control-Allow-Origin", "*")

	// 验证是否来自合法域名访问
	if !l.IsAllowHost() {
		l.Ctx.WriteString(`<!DOCTYPE html><html><head><meta http-equiv="refresh" content="0; url=` +
			website.HostUrl + string([]byte(l.Ctx.Input.Url())[4:]) + `" /></head></html>`)
		l.StopRun()
	}

	l.Data["WebSite"] = website
	l.Data["HostUrl"] = website.HostUrl
	l.Data["WebsiteName"] = website.Title
	l.SetDescript(website.Description)
	l.AddKeyWord(website.Keywords...)

	l.Data["Metes"] = ""
	// l.AddCustomCssStyle("//cdn.bootcss.com/bootstrap/3.3.5/css/", "bootstrap.min.css", "bootstrap-theme.min.css")
	// l.AddCustomCssStyle("//cdn.bootcss.com/font-awesome/4.4.0/css/", "font-awesome.min.css")
	//	l.AddCustomCssStyle("http://fonts.useso.com/", "css?family=Open+Sans:300,400,600&subset=latin,latin-ext")

	// l.AddCustomJsScript("//cdn.bootcss.com/jquery/2.1.4/", "jquery.min.js")
	// l.AddCustomJsScript("//cdn.bootcss.com/bootstrap/3.3.5/css/", "bootstrap.min.js")
	// l.AddCustomCssStyle("//cdn.bootcss.com/jquery-migrate/1.2.1/", "jquery-migrate.min.js")
	// l.AddCustomJsScript("//cdn.bootcss.com/wow/1.1.2/", "wow.min.js")

	l.AddCustomCssStyle("http://7xih3t.com1.z0.glb.clouddn.com/",
		"bootstrap.min.css", "bootstrap-theme.min.css", "font-awesome.min.css")
	l.AddCustomJsScript("http://7xih3s.com1.z0.glb.clouddn.com/",
		"jquery-2.1.3.min.js", "bootstrap.min.js", "jquery-migrate.1.2.1.min.js", "wow1.1.2.js")
	l.AddCssStyle("style.css")

	l.Data["Tail"] = `热眼看社会，冷眼看风险。`

	l.Data["IsPaging"] = true
	l.Data["LastPage"] = 1
	l.Data["Pages"] = []int{1}

	l.LayoutSections = make(map[string]string)

	l.Layout = "layout/index.html"
}
