package models

import (
	"errors"
	logpkg "github.com/fxgcj/website/lib/log"
	"github.com/fxgcj/website/lib/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var (
	E_NOT_OBJ_ID = errors.New("not an bson object id")
)
var (
	log = logpkg.GetLogger()
)

func GetAllBlogs() (blogs Blogs) {
	err := mongo.DB().WithC(mongo.C_BLOGS, func(db *mgo.Collection) error {
		return db.Find(nil).Sort("-created").All(&blogs)
	})
	if err != nil {
		log.Error("Find error, ", err)
	}
	return
}

func GetAllTags() (tags Tags) {
	err := mongo.DB().All(mongo.C_TAGS, nil, &tags)
	if err != nil {
		log.Error("Find error, ", err)
	}
	tags.Sort()
	return
}

func GetAllCategories() (tags Tags) {
	err := mongo.DB().All(mongo.C_CATEGORY, nil, &tags)
	if err != nil {
		log.Error("Find error, ", err)
	}
	tags.Sort()
	return
}

func GetLatestBlogs() (blogs Blogs) {
	err := mongo.DB().WithC(mongo.C_BLOGS, func(db *mgo.Collection) error {
		return db.Find(nil).Sort("-created").Limit(5).All(&blogs)
	})
	if err != nil {
		log.Error("find latest blog  error, ", err)
		return
	}
	return
}

func GetAllMonth() []string {
	var m []string
	var ret [](map[string]interface{})
	err := mongo.DB().All(mongo.C_MONTH, nil, &ret)
	if err != nil {
		log.Error(err)
		return m
	}
	for _, v := range ret {
		log.Debug(v)
		if month, ok := v["name"]; ok {
			m = append(m, month.(string))
		}
	}
	return m
}

//
func GetBlogsGroup(groupType string, name string) (blogs Blogs) {

	switch groupType {
	case "tag":
		err := mongo.DB().WithC(mongo.C_BLOGS, func(db *mgo.Collection) error {
			return db.Find(bson.M{"tags": bson.M{"$all": []string{name}}}).Sort("-created").All(&blogs)
		})
		if err != nil {
			log.Error("Find error, ", err)
			return
		}
		log.Debug("get tag, ", blogs)
		return
	case "category":
		err := mongo.DB().WithC(mongo.C_BLOGS, func(db *mgo.Collection) error {
			return db.Find(bson.M{"category": bson.M{"$all": []string{name}}}).Sort("-created").All(&blogs)
		})
		if err != nil {
			log.Error("Find error, ", err)
			return
		}
		log.Debug("get category, ", blogs)
		return
	default:
		log.Error("GetBlogsByTag type error")
		return
	}
	return
}

func GetMonthBlogs(year, month int) (blogs Blogs) {
	begin := time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.Local)
	end := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.Local)
	err := mongo.DB().WithC(mongo.C_BLOGS, func(db *mgo.Collection) error {
		return db.Find(bson.M{"created": bson.M{"$gte": begin, "$lt": end}}).Sort("-created").All(&blogs)
	})
	if err != nil {
		log.Error("find month error, ", err)
		return
	}
	log.Debug("get month, ", blogs)
	return
}

func SearchBlogs(keyword string) (blogs Blogs) {
	err := mongo.DB().WithC(mongo.C_BLOGS, func(db *mgo.Collection) error {
		return db.Find(
			bson.M{
				"$or": []bson.M{bson.M{"source": bson.M{"$regex": keyword}},
					bson.M{"title": bson.M{"$regex": keyword}},
					bson.M{"summary": bson.M{"$regex": keyword}}},
			}).Sort("-created").All(&blogs)
	})
	if err != nil {
		log.Error("find month error, ", err)
		return
	}
	return
}
