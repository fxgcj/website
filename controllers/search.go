package controllers

import (
	. "github.com/fxgcj/website/models"
)

type SearchController struct {
	ListController
}

func (s *SearchController) Get() {
	word := s.GetString("word")
	page, _ := s.GetInt("page")
	if page > 0 {
		page--
	}
	var begin, end int
	blogs := GetBlogsGroup("category", word)
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

	s.SetPageTitle(word)
	s.AddKeyWord(word)
	s.Data["Blogs"] = blogs[begin:end]
	s.Data["LastestBlogs"] = blogs[:]
	s.Data["Tags"] = GetAllTags()
	s.Data["Category"] = GetAllCategories()
	s.Data["MonthBlog"] = GetAllMonth()

	s.LayoutSections["Sidebar"] = "sidebar.tpl"
	s.TplNames = "list.tpl"
}
