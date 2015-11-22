package controllers

import (
	"github.com/ckeyer/commons/lib"
	"net/http"

	. "github.com/fxgcj/website/models"
)

type AdminController struct {
	BaseController
}

// (b *BaseController)GetBlogs ...
func (c *AdminController) Get() {
	page, _ := c.GetInt("page")
	if page > 0 {
		page--
	}
	var begin, end int
	blogs := GetAllBlogs()
	if count := len(blogs); count > (page+1)*PAGE_STEP {
		begin = page * PAGE_STEP
		end = begin + PAGE_STEP
	} else if count < page*PAGE_STEP {
		begin = (count / PAGE_STEP) * PAGE_STEP
		end = count
	} else {
		begin = page * PAGE_STEP
		end = count
	}

	c.Data["Blogs"] = blogs[begin:end]
	c.Data["LastestBlogs"] = blogs[:]
	c.Data["Tags"] = GetAllTags()
	c.Data["Category"] = GetAllCategories()
	c.Data["MonthBlog"] = blogs.GetMonthSlice()

	c.Layout = "layout/admin.html"
	c.TplNames = "admin/list.tpl"
}

func (a *AdminController) Create() {
	a.AddJsScript("md5.js", "edit.js")

	key_a := lib.RandomInt(5, 49)
	a.SetSession("a", key_a)
	a.Data["a"] = key_a
	key_b := lib.RandomInt(5, 50)
	a.SetSession("b", key_b)
	a.Data["b"] = key_b

	a.Layout = "layout/admin.html"
	a.TplNames = "admin/create.tpl"
}

func (a *AdminController) Post() {
	var blog Blog
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
	id := a.GetString("id")
	blog := GetBlogID(id)
	if blog == nil {
		a.Error(http.StatusNotFound, "not found")
	}
	a.AddJsScript("md5.js", "edit.js")

	key_a := lib.RandomInt(5, 49)
	a.SetSession("a", key_a)
	a.Data["a"] = key_a
	key_b := lib.RandomInt(5, 50)
	a.SetSession("b", key_b)
	a.Data["b"] = key_b

	a.Data["Blog"] = blog

	a.Layout = "layout/admin.html"
	a.TplNames = "admin/edit.tpl"
}

func (a *AdminController) Patch() {
	var blog Blog
	if err := a.ParseForm(&blog); err != nil {
		log.Error(err)
	}
	log.Debug(blog)

	id := a.GetString("id")
	err := blog.UpdateID(id)
	if err != nil {
		a.Error(http.StatusBadRequest, err)
	}
	a.WriteMsg("update successful")
}

func (a *AdminController) Delete() {
	id := a.GetString("id")
	err := DeleteBlogID(id)
	if err != nil {
		a.Error(500, err)
	}
	a.WriteMsg("deleted successful")
}
