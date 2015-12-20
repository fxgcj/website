package controllers

import (
	. "github.com/fxgcj/website/models"
)

type TagController struct {
	ListController
}

func (t *TagController) Get() {
	name := t.GetString("name")
	blogs := GetBlogsGroup("tag", name)

	t.setPaging(len(blogs), PAGE_STEP)

	begin, end := t.getPageStartEnd(len(blogs), PAGE_STEP)

	t.SetPageTitle(name)
	t.AddKeyWord(name)
	t.Data["Blogs"] = blogs[begin:end]
	t.Data["LastestBlogs"] = GetLatestBlogs()
	t.Data["Tags"] = GetAllTags()
	t.Data["Category"] = GetAllCategories()
	t.Data["MonthBlog"] = GetAllMonth()

	t.LayoutSections["Sidebar"] = "sidebar.tpl"
	t.TplNames = "list.tpl"
}
