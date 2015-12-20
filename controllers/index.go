package controllers

import (
	. "github.com/fxgcj/website/models"
)

type IndexController struct {
	ListController
}

func (c *IndexController) Get() {
	blogs := GetAllBlogs()

	c.setPaging(len(blogs), PAGE_STEP)

	begin, end := c.getPageStartEnd(len(blogs), PAGE_STEP)

	c.Data["Blogs"] = blogs[begin:end]
	c.Data["LastestBlogs"] = GetLatestBlogs()
	c.Data["Tags"] = GetAllTags()
	c.Data["Category"] = GetAllCategories()
	c.Data["MonthBlog"] = GetAllMonth()

	c.LayoutSections["Sidebar"] = "sidebar.tpl"

	c.TplNames = "list.tpl"
}
