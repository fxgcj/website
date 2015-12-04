package models

import (
	"fmt"
	"time"

	// "github.com/fxgcj/website/lib/markdown"
	"github.com/fxgcj/website/lib/mgodb"
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
	Content []byte `json:"content" bson:"content" form:"-"`
}

func GetBlogID(id string) (b *Blog) {
	db := mgodb.GetMongoDB()
	b = new(Blog)
	err := db.C(mgodb.C_BLOGS).FindId(bson.ObjectIdHex(id)).One(b)
	if err != nil {
		log.Error("get blog error, ", err)
	}
	//	log.Debugf("find blog , %#v", b)
	return
}

func DeleteBlogID(id string) (err error) {
	blog := &Blog{ID: bson.ObjectIdHex(id)}
	return blog.Delete()
}

func (b *Blog) Insert() (err error) {
	b.ID = bson.NewObjectId()
	b.Created = time.Now()
	b.Updated = time.Now()
	b.Author = "风险观察君"
	if len(b.Source) > 0 {
		b.Content = []byte(b.Source) //markdown.Trans2html([]byte(b.Source))
	}
	db := mgodb.GetMongoDB()
	err = db.C(mgodb.C_BLOGS).Insert(b)
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
	db := mgodb.GetMongoDB()
	err = db.C(mgodb.C_BLOGS).FindId(bid).One(old)
	if err != nil {
		return
	}

	b.ID = old.ID
	b.Created = old.Created
	b.Author = old.Author
	b.Updated = time.Now()

	if len(b.Source) > 0 {
		b.Content = []byte(b.Source) //markdown.Trans2html([]byte(b.Source))
	}
	err = db.C(mgodb.C_BLOGS).UpdateId(b.ID, b)
	if err != nil {
		return err
	}
	_, err = b.removeAllGroup()
	if err != nil {
		return err
	}
	err = b.insertGroup()
	return
}

func (b *Blog) Delete() (err error) {
	db := mgodb.GetMongoDB()
	err = db.C(mgodb.C_BLOGS).RemoveId(b.ID)
	if err != nil {
		return err
	}
	_, err = b.removeAllGroup()
	return err
}

// 插入文章的tag信息到数据库
func (b *Blog) insertGroup() error {
	db := mgodb.GetMongoDB()
	for _, tag := range b.Tags {
		if tag == "" {
			continue
		}
		info, err := db.C(mgodb.C_TAGS).Upsert(bson.M{"name": tag},
			bson.M{"$push": bson.M{"blogs": b.ID}})
		if err != nil {
			return fmt.Errorf("%s %s\n", info, err)
		}
	}
	for _, tag := range b.Category {
		if tag == "" {
			continue
		}
		info, err := db.C(mgodb.C_CATEGORY).Upsert(bson.M{"name": tag},
			bson.M{"$push": bson.M{"blogs": b.ID}})
		if err != nil {
			return fmt.Errorf("%s %s\n", info, err)
		}
	}

	_, err := db.C(mgodb.C_MONTH).Upsert(bson.M{"name": fmt.Sprintf("%d-%d", b.Created.Year(), b.Created.Month())},
		bson.M{"$push": bson.M{"blogs": b.ID}})

	return err
}

// 移除文章所有标签
func (b *Blog) removeAllGroup() (count int, err error) {
	db := mgodb.GetMongoDB()
	info, err := db.C(mgodb.C_TAGS).Upsert(nil,
		bson.M{"$pull": bson.M{"blogs": b.ID}})
	if err != nil {
		return
	}
	count = info.Removed

	info, err = db.C(mgodb.C_CATEGORY).Upsert(nil,
		bson.M{"$pull": bson.M{"blogs": b.ID}})
	if err != nil {
		return
	}
	count += info.Removed

	info, err = db.C(mgodb.C_MONTH).Upsert(nil, bson.M{"$pull": bson.M{"blogs": b.ID}})
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
