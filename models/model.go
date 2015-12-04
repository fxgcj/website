package models

import (
	"errors"
	logpkg "github.com/fxgcj/website/lib/log"
	"github.com/fxgcj/website/lib/mgodb"
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
	db := mgodb.GetMongoDB()
	err := db.C(mgodb.C_BLOGS).Find(nil).Sort("-created").All(&blogs)
	if err != nil {
		log.Error("Find error, ", err)
	}
	return
}

func GetAllTags() (tags Tags) {
	db := mgodb.GetMongoDB()
	err := db.C(mgodb.C_TAGS).Find(nil).All(&tags)
	if err != nil {
		log.Error("Find error, ", err)
	}
	tags.Sort()
	return
}

func GetAllCategories() (tags Tags) {
	db := mgodb.GetMongoDB()
	err := db.C(mgodb.C_CATEGORY).Find(nil).All(&tags)
	if err != nil {
		log.Error("Find error, ", err)
	}
	tags.Sort()
	return
}

func GetAllMonth() []string {
	var m []string
	var ret [](map[string]interface{})
	db := mgodb.GetMongoDB()
	err := db.C(mgodb.C_MONTH).Find(nil).All(&ret)
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
	db := mgodb.GetMongoDB()

	switch groupType {
	case "tag":
		err := db.C(mgodb.C_BLOGS).Find(bson.M{"tags": bson.M{"$all": []string{name}}}).Sort("-created").All(&blogs)
		if err != nil {
			log.Error("Find error, ", err)
			return
		}
		log.Debug("get tag, ", blogs)
		return
	case "category":
		err := db.C(mgodb.C_BLOGS).Find(bson.M{"category": bson.M{"$all": []string{name}}}).Sort("-created").All(&blogs)
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
	db := mgodb.GetMongoDB()
	begin := time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.Local)
	end := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.Local)
	err := db.C(mgodb.C_BLOGS).Find(bson.M{"created": bson.M{"$gte": begin, "$lt": end}}).Sort("-created").All(&blogs)
	if err != nil {
		log.Error("find month error, ", err)
		return
	}
	log.Debug("get month, ", blogs)
	return
}

func SearchBlogs(keyword string) (blogs Blogs) {
	db := mgodb.GetMongoDB()
	err := db.C(mgodb.C_BLOGS).Find(
		bson.M{
			"$or": []bson.M{bson.M{"source": bson.M{"$regex": keyword}},
				bson.M{"title": bson.M{"$regex": keyword}},
				bson.M{"summary": bson.M{"$regex": keyword}}},
		}).Sort("-created").All(&blogs)
	if err != nil {
		log.Error("find month error, ", err)
		return
	}
	return
}
