package controllers

import (
	. "github.com/fxgcj/website/models"
)

type BlogController struct {
	BaseController
}

// (b *BaseController)GetBlogs ...
func (b *BlogController) Get() {

	id := b.Ctx.Input.Param(":id")
	blog := GetBlogID(id)
	if blog == nil {
		log.Debug("id: ", id)
		return
	}
	b.SetPageTitle(blog.Title)
	b.AddKeyWord(blog.Tags...)
	b.AddKeyWord(blog.Category...)
	b.SetDescript(blog.Summary)
	b.Data["Blog"] = blog
	b.Data["BContent"] = blog.Content
	b.Data["DS_key"] = blog.ID.Hex()
	log.Debug("DS_key: ", b.Data["DS_key"])
	b.Data["DS_title"] = blog.Title

	blogs := GetAllBlogs()
	b.Data["LastestBlogs"] = blogs[:]
	b.Data["Tags"] = GetAllTags()
	b.Data["Category"] = GetAllCategories()
	b.Data["MonthBlog"] = GetAllMonth()

	b.LayoutSections["Sidebar"] = "sidebar.tpl"
	b.TplNames = "show.tpl"
}
