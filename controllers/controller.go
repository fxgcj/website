package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/ckeyer/commons/lib"
	"github.com/fxgcj/website/conf"
	logpkg "github.com/fxgcj/website/lib/log"
	"net/http"
	"strings"
)

const (
	PAGE_STEP = 2

	COOKIE_SECRET     = "cookie_secret"
	COOKIE_SECRET_LEN = 15
)

var (
	website *conf.WebSite
	log     = logpkg.GetLogger()
)

type Controller struct {
	beego.Controller
}

func (c *Controller) WriteJSON(code int, data interface{}) {
	c.Ctx.Output.SetStatus(code)
	bs, err := json.Marshal(data)
	if err != nil {
		log.Error("WriteJSON Error.. Marshal err, ", err)
	}
	c.Ctx.ResponseWriter.Write(bs)
	c.StopRun()
}

func (c *Controller) Error(code int, msg interface{}) {
	data := make(map[string]string)
	data["error"] = fmt.Sprint(msg)
	c.WriteJSON(code, data)
}

func (c *Controller) WriteMsg(msg interface{}) {
	data := make(map[string]string)
	data["msg"] = fmt.Sprint(msg)
	c.WriteJSON(http.StatusOK, data)
}

// 是否是通过允许的域名访问
func (c *Controller) IsAllowHost() bool {
	host := c.Ctx.Input.Host()
	for _, v := range website.EnableDomain {
		if strings.Index(host, v) >= 0 {
			return true
		}
	}
	log.Debugf("Not Enable Domain %s\n", c.Ctx.Input.Host())
	return false
}

//  (b *Controller)SetPateTitle 设置页面显示标题
func (c *Controller) SetPageTitle(title string) {
	c.Data["PageTitle"] = website.Title + " - " + title
}

// (b *Controller)SetKeyWord 设置或更新Meta关键字
func (c *Controller) SetKeyWord(args ...string) {
	c.Data["Keywords"] = strings.Join(args, ",")
}

// (b *Controller)AddKeyWord 添加Meta关键字
func (c *Controller) AddKeyWord(args ...string) {
	if c.Data["Keywords"] != nil {
		keyword := strings.Split(fmt.Sprint(c.Data["Keywords"]), ",")
		c.Data["Keywords"] = strings.Join(append(keyword, args...), ",")
	} else {
		c.SetKeyWord(args...)
	}
}

// get cookie's secret from session or create
func (c *Controller) getCookieSecret() string {
	cscret := c.GetSession(COOKIE_SECRET)
	if cscret != nil && len(cscret.(string)) == COOKIE_SECRET_LEN {
		return cscret.(string)
	} else {
		s := lib.RandomString(COOKIE_SECRET_LEN)
		c.SetSession(COOKIE_SECRET, s)
		return s
	}
}

func (c *Controller) SetCookie(key, value string, others ...interface{}) {
	c.SetSecureCookie(c.getCookieSecret(), key, value, others...)
}

func (c *Controller) GetCookie(key string) (string, bool) {
	return c.GetSecureCookie(c.getCookieSecret(), key)
}

// (b *Controller)SetDescript 设置Mets描述
func (c *Controller) SetDescript(des string) {
	c.Data["Description"] = des
}

// (b *Controller)AddJsScript 按配置的js路径添加js文件
func (c *Controller) AddJsScript(args ...string) {
	c.AddCustomJsScript(website.JsUrl, args...)
}

// (b *Controller)AddCustomJsScript 添加自定义js
func (c *Controller) AddCustomJsScript(src_url string, args ...string) {
	var jstags []string

	if c.Data["Scripts"] != nil {
		jstags = strings.Split(fmt.Sprint(c.Data["Scripts"]), "\n")
	}
	for _, js := range args {
		newtag := fmt.Sprintf(`<script type="text/javascript" src="%s%s"></script>`, src_url, js)
		jstags = append(jstags, newtag)
	}
	c.Data["Scripts"] = strings.Join(jstags, "\n")

}

// (b *Controller)AddCssStyle 按配置的css路径添加css文件
func (c *Controller) AddCssStyle(args ...string) {
	c.AddCustomCssStyle(website.CssUrl, args...)
}

// (b *Controller)AddCustomCssStyle 添加自定义css
func (c *Controller) AddCustomCssStyle(src_url string, args ...string) {
	var csstags []string

	if c.Data["Styles"] != nil {
		csstags = strings.Split(fmt.Sprint(c.Data["Styles"]), "\n")
	}
	for _, css := range args {
		newtag := fmt.Sprintf(`<link rel="stylesheet" media="screen" type="text/css" href="%s%s"/>`, src_url, css)
		csstags = append(csstags, newtag)
	}
	c.Data["Styles"] = strings.Join(csstags, "\n")
}

func (c *Controller) setPaging(count, step int) {
	log.Debugf("%#v", c.Ctx.Input.Params)
	c.Data["IsPaging"] = false
	if count > step {
		lastPage := (count + (step+1)/2) / step

		oldQuery := make(map[string][]string)
		querys := make([]string, 0, lastPage-1)

		if c.Ctx.Input.Request.Form == nil {
			c.Ctx.Request.ParseForm()
		}
		for k, v := range c.Ctx.Request.Form {
			oldQuery[k] = v
		}
		for i := 0; i < lastPage; i++ {
			oldQuery["page"] = []string{fmt.Sprint(i + 1)}
			cells := make([]string, 0, len(oldQuery))
			for k, v := range oldQuery {
				cells = append(cells, fmt.Sprintf("%s=%s", k, strings.Join(v, ",")))
			}
			querys = append(querys, "?"+strings.Join(cells, "&"))
		}
		log.Debugf("querys %s", querys)
		c.Data["IsPaging"] = true
		c.Data["LastPage"] = querys[len(querys)-1]
		c.Data["FirstPage"] = querys[0]
		c.Data["Query"] = querys
	}
}

func (c *Controller) getPageStartEnd(count, step int) (begin, end int) {
	page, _ := c.GetInt("page")
	if page > 0 {
		page--
	}
	if count > (page+1)*step {
		begin = page * step
		end = begin + step
	} else if count < page*step {
		begin = (count / step) * step
		end = count
	} else {
		begin = page * step
		end = count
	}
	return
}
