package controllers

import (
	"fmt"
	. "github.com/fxgcj/website/models"
	"strconv"
)

type ArchiveController struct {
	BaseController
}

func (c *ArchiveController) Get() {

	year, _ := strconv.Atoi(c.Ctx.Input.Param(":year"))
	month, _ := strconv.Atoi(c.Ctx.Input.Param(":month"))

	page, _ := c.GetInt("page")
	if page > 0 {
		page--
	}
	var begin, end int
	blogs := GetMonthBlogs(year, month)
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

	c.SetPageTitle(fmt.Sprintf("%d年%d月", year, month))
	c.Data["Blogs"] = blogs[begin:end]
	c.Data["LastestBlogs"] = blogs[:]
	c.Data["Tags"] = GetAllTags()
	c.Data["Category"] = GetAllCategories()
	c.Data["MonthBlog"] = GetAllMonth()

	c.LayoutSections["Sidebar"] = "sidebar.tpl"
	c.TplNames = "list.tpl"
}
