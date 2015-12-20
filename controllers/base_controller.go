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
			website.HostUrl + string([]byte(b.Ctx.Input.Url())[1:]) + `" /></head></html>`)
		b.StopRun()
	}

	b.Data["WebSite"] = website
	b.Data["WebsiteName"] = website.Title
	b.Data["HostUrl"] = website.HostUrl
	b.Data["Metes"] = `<meta property="qc:admins" content="3360317257660736727" />
	<meta property="wb:webmaster" content="104fa85e4f2b2606" />`
	b.Data["Tail"] = `热眼看社会，冷眼看风险。`
	b.Data["FriendLinks"] = website.FriendLinks
	b.SetPageTitle("首页")
	b.SetDescript(website.Description)
	b.AddKeyWord(website.Keywords...)

	// 七牛库
	b.AddCustomCssStyle("http://7xih3t.com1.z0.glb.clouddn.com/",
		"bootstrap.min.css", "bootstrap-theme.min.css", "font-awesome.min.css")
	b.AddCustomJsScript("http://7xih3s.com1.z0.glb.clouddn.com/",
		"jquery-2.1.3.min.js", "bootstrap.min.js", "jquery-migrate.1.2.1.min.js", "wow1.1.2.js")

	// bootstrap CDN源
	// b.AddCustomCssStyle("//cdn.bootcss.com/bootstrap/3.3.5/css/", "bootstrap.min.css", "bootstrap-theme.min.css")
	// b.AddCustomCssStyle("//cdn.bootcss.com/font-awesome/4.4.0/css/", "font-awesome.min.css")

	// b.AddCustomJsScript("//cdn.bootcss.com/jquery/2.1.4/", "jquery.min.js")
	// b.AddCustomJsScript("//cdn.bootcss.com/bootstrap/3.3.5/js/", "bootstrap.min.js")
	// b.AddCustomJsScript("//cdn.bootcss.com/jquery-migrate/1.2.1/", "jquery-migrate.min.js")
	// b.AddCustomJsScript("//cdn.bootcss.com/wow/1.1.2/", "wow.min.js")
	b.AddCssStyle("style.css")

	b.Data["IsPaging"] = false
	b.Data["LastPage"] = 1
	b.Data["Pages"] = []int{1}

	b.LayoutSections = make(map[string]string)

	b.Layout = "layout/index.html"
}
