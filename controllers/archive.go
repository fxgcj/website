package controllers

import (
	"fmt"
	. "github.com/fxgcj/website/models"
	"strconv"
)

type ArchiveController struct {
	ListController
}

func (a *ArchiveController) Get() {

	year, _ := strconv.Atoi(a.Ctx.Input.Param(":year"))
	month, _ := strconv.Atoi(a.Ctx.Input.Param(":month"))

	blogs := GetMonthBlogs(year, month)

	a.setPaging(len(blogs), PAGE_STEP)

	begin, end := a.getPageStartEnd(len(blogs), PAGE_STEP)

	a.SetPageTitle(fmt.Sprintf("%d年%d月", year, month))
	a.Data["Blogs"] = blogs[begin:end]
	a.Data["LastestBlogs"] = GetLatestBlogs()
	a.Data["Tags"] = GetAllTags()
	a.Data["Category"] = GetAllCategories()
	a.Data["MonthBlog"] = GetAllMonth()

	a.LayoutSections["Sidebar"] = "sidebar.tpl"
	a.TplNames = "list.tpl"
}
