package controllers

import (
	. "github.com/fxgcj/website/models"
)

type SearchController struct {
	ListController
}

func (s *SearchController) Get() {
	keyword := s.GetString("keyword")

	blogs := SearchBlogs(keyword)

	s.setPaging(len(blogs), PAGE_STEP)

	begin, end := s.getPageStartEnd(len(blogs), PAGE_STEP)

	s.SetPageTitle(keyword)
	s.AddKeyWord(keyword)
	s.Data["Blogs"] = blogs[begin:end]
	s.Data["LastestBlogs"] = GetLatestBlogs()
	s.Data["Tags"] = GetAllTags()
	s.Data["Category"] = GetAllCategories()
	s.Data["MonthBlog"] = GetAllMonth()

	s.LayoutSections["Sidebar"] = "sidebar.tpl"
	s.TplNames = "list.tpl"
}
