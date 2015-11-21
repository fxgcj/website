package controllers

import (
	"github.com/ckeyer/commons/lib"
	"net/http"

	"github.com/fxgcj/website/models"
)

type AdminController struct {
	BaseController
}

// (b *BaseController)GetBlogs ...
func (a *AdminController) Get() {
	// b.Data["LastestBlogs"] = models.GetBlogs(0, 5)
	// b.Data["Tags"] = models.GetAllTags()
	// b.Data["Category"] = models.GetAllCategory()
	// b.Data["MonthBlog"] = models.GetAllMonth()

	// name := b.Ctx.Input.Param(":name")
	// blog := models.GetBlog(name)
	// if blog == nil {
	// 	log.Debug("name: ", name)
	// 	return
	// }
	// b.Data["Blog"] = blog
	// b.Data["BContent"] = string(blog.Content)
	log.Debug("fuck")
	// b.LayoutSections["Sidebar"] = "sidebar.tpl"
	// b.LayoutSections["Duoshuo"] = "duoshuo.tpl"
	a.Data["fuck"] = "fuck"

	a.Layout = "layout/admin.html"
	a.TplNames = "admin/list.tpl"
}

func (a *AdminController) Create() {
	a.AddJsScript("md5.js", "edit.js")
	a.Data["fuck"] = "fuck"

	key_a := lib.RandomInt(5, 49)
	a.SetSession("a", key_a)
	a.Data["a"] = key_a
	key_b := lib.RandomInt(5, 50)
	a.SetSession("b", key_b)
	a.Data["b"] = key_b

	a.Layout = "layout/admin.html"
	a.TplNames = "admin/edit.tpl"
}

func (a *AdminController) Post() {
	var blog models.Blog
	if err := a.ParseForm(&blog); err != nil {
		log.Error(err)
	}
	blog.AuthorEndpoint = a.Ctx.Input.IP()
	err := blog.Insert()
	if err != nil {
		log.Error("insert failed", err)
		a.Error(http.StatusBadRequest, err)
	}
	log.Notice("Inserted a blog successful")
	a.WriteMsg("inserted successful")
}

func (a *AdminController) Edit() {
	a.AddJsScript("md5.js", "edit.js")
	a.Data["fuck"] = "fuck"

	key_a := lib.RandomInt(5, 49)
	a.SetSession("a", key_a)
	a.Data["a"] = key_a
	key_b := lib.RandomInt(5, 50)
	a.SetSession("b", key_b)
	a.Data["b"] = key_b

	a.Layout = "layout/admin.html"
	a.TplNames = "admin/edit.tpl"
}

func (a *AdminController) Patch() {
	var blog models.Blog
	if err := a.ParseForm(&blog); err != nil {
		log.Error(err)
	}
	a.Ctx.Request.ParseForm()
	log.Debug(blog)

	for k, v := range a.Ctx.Request.Form {
		log.Debug(k, "...", v)
	}
	a.Ctx.WriteString("content")
}
