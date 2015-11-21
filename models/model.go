package models

import (
	logpkg "github.com/fxgcj/website/lib/log"
	"github.com/fxgcj/website/lib/mgodb"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var (
	log = logpkg.GetLogger()
)

func GetAllBlogs() (blogs Blogs) {
	db := mgodb.GetMongoDB()
	err := db.C(mgodb.C_BLOGS).Find(nil).Sort("-updated").All(&blogs)
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

func GetBlogsMonth(year, month int) (blogs Blogs) {
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
