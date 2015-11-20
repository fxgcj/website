package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Blog struct {
	ID       bson.ObjectId `json:"id" bson:"_id" form:"-"`
	Name     string        `json:"name" bson:"name" form:"-"`
	Author   string        `json:"author" bson:"author" form:"-"`
	Head     string        `json:"head" bson:"head" form:"-"`
	Created  time.Time     `json:"created" bson:"created" form:"-"`
	Title    string        `json:"title" bson:"title" form:"title"`
	Tags     []string      `json:"tags" bson:"tags" form:"tags[]"`
	Category []string      `json:"category" bson:"category" form:"category[]"`
	Status   string        `json:"status" bson:"status form:"-"`
	Summary  string        `json:"summary" bson:"summary" form:"summary"`

	Secret string `json:"-" bson:"-" form:"secret"`
	// .md 源文件
	Source string `json:"source" bson:"source" form:"content"`
	// .html 文件
	Content string `json:"content" bson:"content" form:"-"`
}

// // GetBlogs
// func GetBlogs(start, count int) (bs []*Blog) {
// 	bs = MyPool.GetBlogs(start, count)
// 	return
// }

// // GetBlog
// func GetBlog(name string) *Blog {
// 	b := MyPool.GetBlogByName(name)
// 	return b
// }

// // GetCount
// func GetCount() int {
// 	return MyPool.GetCount()
// }
