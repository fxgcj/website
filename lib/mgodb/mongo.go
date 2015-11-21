package mgodb

import (
	"fmt"
	"github.com/fxgcj/website/conf"
	"gopkg.in/mgo.v2"
)

var mongoConfig *conf.MongoDB

const (
	DB_ARTICLE = "fxgcj"

	C_BLOGS    = "blogs"
	C_TAGS     = "tags"
	C_CATEGORY = "categories"
	C_MONTH    = "month"
)

var db *mgo.Database

func GetMongoDB() *mgo.Database {
	if db != nil {
		return db
	}
	mconf := conf.GetConf().MongoDB
	mgo_conn_url := fmt.Sprintf("%s:%d", mconf.Host, mconf.Port)
	session, err := mgo.Dial(mgo_conn_url)
	if err != nil {
		panic(err)
	}
	return session.DB(DB_ARTICLE)
}
