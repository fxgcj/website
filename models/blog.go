package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Blog struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Author   string        `json:"author" bson:"author"`
	Head     string        `json:"head" bson:"head"`
	Created  time.Time     `json:"created" bson:"created"`
	Title    string        `json:"title" bson:"title"`
	Tags     []string      `json:"tags" bson:"tags"`
	Category []string      `json:"category" bson:"category"`
	Status   string        `json:"status" bson:"status"`
	Summary  string        `json:"summary" bson:"summary"`

	// .md 源文件
	Source string `json:"source" bson:"source"`
	// .html 文件
	Content string `json:"content" bson:"content"`
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
