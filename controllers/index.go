package controllers

import (
	. "github.com/fxgcj/website/models"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Get() {
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

	c.LayoutSections["Sidebar"] = "sidebar.tpl"

	c.TplNames = "list.tpl"
}
