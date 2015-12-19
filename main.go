package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/fxgcj/website/conf"
	logpkg "github.com/fxgcj/website/lib/log"
	"github.com/fxgcj/website/models"
	"github.com/fxgcj/website/routers"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strings"
	"time"
)

var log = logpkg.GetLogger()
var config *conf.Config

func init() {
	var err error
	config, err = conf.LoadConf("conf/v2.json")
	if err != nil {
		panic(err)
	}
	BeegoInit()
}

func main() {
	pushBaidu()
	beego.Run()
}

func BeegoInit() {
	beego.RunMode = config.AppConfig.RunMode
	beego.AppName = config.AppConfig.Name
	beego.HttpPort = config.AppConfig.Port
	beego.BeegoServerName = config.AppConfig.ServerName
	beego.SessionOn = true
	beego.SessionName = "ckeyer"
	// beego.SessionDomain = "fxgcj.org"
	beego.SessionAutoSetCookie = true
	beego.AddFuncMap("showMonth", func(m int) string {
		mstr := fmt.Sprint(m)
		if len(mstr) == 6 {
			return fmt.Sprintf("%s-%s", mstr[:4], mstr[4:])
		}
		return mstr
	})
	beego.AddFuncMap("showDate", func(t time.Time) string {
		return fmt.Sprintf("%d-%02d-%02d %02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
	})
	beego.AddFuncMap("showObjectID", func(obj bson.ObjectId) string {
		return obj.Hex()
	})
	beego.AddFuncMap("setURLMonth", func(m string) string {
		mstr := strings.Split(m, "-")
		if len(mstr) == 2 {
			return fmt.Sprintf("%s/%s", mstr[0], mstr[1])
		}
		return ""
	})

	routers.LoadRouters()
}

func pushBaidu() {
	go func() {
		for ; ; time.Sleep(time.Hour * 2) {
			urls := models.GetAllURLs()
			if len(urls) < 2 {
				log.Warning("get all urls : ", len(urls))
				continue
			}
			api := `http://data.zz.baidu.com/urls?site=www.fxgc.org&token=qkOcqAbvlHmU8ukC&type=original`
			bodyType := "text/plain"
			body := strings.NewReader(strings.Join(urls, "\n"))
			req, err := http.Post(api, bodyType, body)
			if err != nil || req.StatusCode > 200 {
				log.Error("push baidu err, ", err)
			}
			log.Notice("push baidu successful")
		}
	}()
}
