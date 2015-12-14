package models

import (
	"fmt"
	"time"

	// "github.com/fxgcj/website/lib/markdown"
	"github.com/fxgcj/website/lib/mongo"
	"gopkg.in/mgo.v2/bson"
)

type Blog struct {
	ID             bson.ObjectId `json:"id" bson:"_id" form:"-"`
	Name           string        `json:"name" bson:"name" form:"-"`
	Author         string        `json:"author" bson:"author" form:"-"`
	AuthorEndpoint string        `json:"-" bson:"author_endpoint" form:"-"`
	Head           string        `json:"head" bson:"head" form:"-"`
	Created        time.Time     `json:"created" bson:"created" form:"-"`
	Updated        time.Time     `json:"updated" bson:"updated" form:"-"`
	Title          string        `json:"title" bson:"title" form:"title"`
	Tags           []string      `json:"tags" bson:"tags" form:"tags[]"`
	Category       []string      `json:"category" bson:"category" form:"category[]"`
	Status         string        `json:"status" bson:"status form:"-"`
	Summary        string        `json:"summary" bson:"summary" form:"summary"`
	Link           string        `json:"link" bson:"link" form:"link"`
	Views          int           `json:"views" bson:"views" form:"-"`

	Secret string `json:"-" bson:"-" form:"secret"`
	// .md 源文件
	Source string `json:"source" bson:"source" form:"content"`
	// .html 文件
	Content string `json:"content" bson:"content" form:"-"`
}

func GetBlogID(id string) (b *Blog) {
	if !bson.IsObjectIdHex(id) {
		return
	}

	b = new(Blog)
	err := mongo.DB().FindId(mongo.C_BLOGS, bson.ObjectIdHex(id), b)
	if err != nil {
		log.Error("get blog error, ", err)
	}
	//	log.Debugf("find blog , %#v", b)
	return
}

func DeleteBlogID(id string) (err error) {
	if !bson.IsObjectIdHex(id) {
		return E_NOT_OBJ_ID
	}

	blog := &Blog{ID: bson.ObjectIdHex(id)}
	return blog.Delete()
}

func (b *Blog) Insert() (err error) {
	b.ID = bson.NewObjectId()
	b.Created = time.Now()
	b.Updated = time.Now()
	b.Author = "风险观察君"
	if len(b.Source) > 0 {
		b.Content = b.Source //markdown.Trans2html([]byte(b.Source))
	}
	err = mongo.DB().Insert(mongo.C_BLOGS, b)
	if err != nil {
		return err
	}
	err = b.insertGroup()
	return
}

func (b *Blog) UpdateID(id string) (err error) {
	if !bson.IsObjectIdHex(id) {
		return E_NOT_OBJ_ID
	}
	bid := bson.ObjectIdHex(id)
	old := new(Blog)
	err = mongo.DB().FindId(mongo.C_BLOGS, bid, old)
	if err != nil {
		return
	}

	b.ID = old.ID
	b.Created = old.Created
	b.Author = old.Author
	b.Updated = time.Now()

	if len(b.Source) > 0 {
		b.Content = b.Source //markdown.Trans2html([]byte(b.Source))
	}
	err = mongo.DB().UpdateId(mongo.C_BLOGS, b.ID, b)
	if err != nil {
		return err
	}
	err = b.removeAllGroup()
	if err != nil {
		return err
	}
	err = b.insertGroup()
	return
}

func (b *Blog) Delete() (err error) {
	err = mongo.DB().RemoveId(mongo.C_BLOGS, b.ID)
	if err != nil {
		return err
	}
	err = b.removeAllGroup()
	return err
}

// 插入文章的tag信息到数据库
func (b *Blog) insertGroup() error {
	for _, tag := range b.Tags {
		if tag == "" {
			continue
		}
		err := mongo.DB().Upsert(mongo.C_TAGS, bson.M{"name": tag}, bson.M{"$push": bson.M{"blogs": b.ID}})
		if err != nil {
			return fmt.Errorf("%s\n", err)
		}
	}
	for _, tag := range b.Category {
		if tag == "" {
			continue
		}
		err := mongo.DB().Upsert(mongo.C_CATEGORY, bson.M{"name": tag},
			bson.M{"$push": bson.M{"blogs": b.ID}})
		if err != nil {
			return fmt.Errorf("%s\n", err)
		}
	}

	err := mongo.DB().Upsert(mongo.C_MONTH, bson.M{"name": fmt.Sprintf("%d-%d", b.Created.Year(), b.Created.Month())},
		bson.M{"$push": bson.M{"blogs": b.ID}})

	return err
}

// 移除文章所有标签
func (b *Blog) removeAllGroup() (err error) {
	err = mongo.DB().UpdateAll(mongo.C_TAGS, nil, bson.M{"$pull": bson.M{"blogs": b.ID}})
	if err != nil {
		return
	}

	err = mongo.DB().UpdateAll(mongo.C_CATEGORY, nil, bson.M{"$pull": bson.M{"blogs": b.ID}})
	if err != nil {
		return
	}

	err = mongo.DB().UpdateAll(mongo.C_MONTH, nil, bson.M{"$pull": bson.M{"blogs": b.ID}})
	return
}

type Blogs []*Blog

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
