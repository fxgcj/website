package controllers

import (
	"net/http"

	. "github.com/fxgcj/website/models"
)

type AdminController struct {
	AdminBaseController
}

func (a *AdminController) Login() {
	switch a.Ctx.Input.Method() {
	case "POST":
		a.postLogin()
	case "GET":
		a.getLogin()
	default:
		a.Error(http.StatusMethodNotAllowed, "ali")
	}
}
func (a *AdminController) postLogin() {
	var form struct {
		Passowrd string `form:"password"`
		Remember bool   `form:"remember"`
	}
	a.ParseForm(&form)
	log.Debug(form)

	if a.verifyA_B(form.Passowrd) {
		if form.Remember {
			a.SetCookie(COOKIE_IS_LOGINED, "true", 60*60*24*365)
		} else {
			a.SetCookie(COOKIE_IS_LOGINED, "true")
		}
		a.WriteMsg("login successful")
	}
	a.Error(http.StatusBadRequest, "login error")
}
func (a *AdminController) getLogin() {
	a.AddJsScript("md5.js", "edit.js")
	a.setA_B()
	a.TplNames = "admin/login.tpl"
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
	c.Data["MonthBlog"] = GetAllMonth()

	c.Layout = "layout/admin.html"
	c.TplNames = "admin/list.tpl"
}

func (a *AdminController) Create() {
	a.AddJsScript("md5.js", "edit.js")

	a.setA_B()

	a.Layout = "layout/admin.html"
	a.TplNames = "admin/create.tpl"
}

func (a *AdminController) Post() {
	var blog Blog
	if err := a.ParseForm(&blog); err != nil {
		log.Error(err)
	}
	if !a.verifyA_B(blog.Secret) {
		a.Error(404, "auth error")
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
		log.Error("not found")
		a.Error(http.StatusNotFound, "not found")
	}
	a.AddJsScript("md5.js", "edit.js")

	a.setA_B()
	a.Data["Blog"] = blog

	a.Layout = "layout/admin.html"
	a.TplNames = "admin/edit.tpl"
}

func (a *AdminController) Patch() {
	var blog Blog
	if err := a.ParseForm(&blog); err != nil {
		log.Error(err)
	}
	if !a.verifyA_B(blog.Secret) {
		a.Error(404, "auth error")
	}

	id := a.GetString("id")
	err := blog.UpdateID(id)
	if err != nil {
		log.Error(err)
		a.Error(http.StatusBadRequest, err)
	}
	a.WriteMsg("update successful")
}

func (a *AdminController) Delete() {
	id := a.GetString("id")
	sec := a.GetString("secret")
	if !a.verifyA_B(sec) {
		a.Error(404, "auth error")
	}
	err := DeleteBlogID(id)
	if err != nil {
		a.Error(404, err)
	}
	a.WriteMsg("deleted successful")
}
