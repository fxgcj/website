package controllers

import (
	. "github.com/fxgcj/website/models"
)

type CategoryController struct {
	ListController
}

func (c *CategoryController) Get() {
	name := c.GetString("name")
	blogs := GetBlogsGroup("category", name)

	c.setPaging(len(blogs), PAGE_STEP)

	begin, end := c.getPageStartEnd(len(blogs), PAGE_STEP)

	c.SetPageTitle(name)
	c.AddKeyWord(name)
	c.Data["Blogs"] = blogs[begin:end]
	c.Data["LastestBlogs"] = GetLatestBlogs()
	c.Data["Tags"] = GetAllTags()
	c.Data["Category"] = GetAllCategories()
	c.Data["MonthBlog"] = GetAllMonth()

	c.LayoutSections["Sidebar"] = "sidebar.tpl"
	c.TplNames = "list.tpl"
}
