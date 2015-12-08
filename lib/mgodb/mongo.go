package mgodb

import (
	"fmt"
	"github.com/fxgcj/website/conf"
	"github.com/fxgcj/website/lib/audit"
	logpkg "github.com/fxgcj/website/lib/log"
	"gopkg.in/mgo.v2"
	"time"
)

const (
	DB_ARTICLE = "fxgcj"

	C_BLOGS    = "blogs"
	C_TAGS     = "tags"
	C_CATEGORY = "categories"
	C_MONTH    = "group_month"
)

var (
	session      *mgo.Session
	db           *mgo.Database
	mongoConfig  *conf.MongoDB
	mgo_conn_url string
	log          = logpkg.GetLogger()
)

func init() {
	mconf := conf.GetConf().MongoDB
	mgo_conn_url = fmt.Sprintf("%s:%d", mconf.Host, mconf.Port)
}

func GetMongoDB() *mgo.Database {
	var err error
	if session != nil && db != nil {
		if err := session.Ping(); err == nil {
			return db
		}
	}

RECONNECT:
	session, db = nil, nil
	session, err = mgo.Dial(mgo_conn_url)
	if err != nil {
		log.Error(err)
		audit.AuditError(mgo_conn_url, err)
		time.Sleep(time.Second * 3)
		goto RECONNECT
	}
	db = session.DB(DB_ARTICLE)
	return db
}
