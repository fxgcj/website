package models

import (
	"fmt"
	"github.com/fxgcj/website/conf"
	"strings"
)

func GetAllURLs() []string {
	var urls []string
	website := conf.GetConf().WebSite
	home := website.HostUrl

	urls = append(urls, home)

	blogs := GetAllBlogs()
	for _, b := range blogs {
		urls = append(urls, fmt.Sprintf("%sblog/%s", home, b.ID.Hex()))
	}

	tags := GetAllTags()
	for _, t := range tags {
		urls = append(urls, fmt.Sprintf("%stag?name=%s", home, t.Name))
	}

	cate := GetAllTags()
	for _, c := range cate {
		urls = append(urls, fmt.Sprintf("%scategory?name=%s", home, c.Name))
	}

	month := GetAllMonth()
	for _, m := range month {
		ym := strings.Split(m, "-")
		if len(ym) != 2 {
			continue
		}
		urls = append(urls, fmt.Sprintf("%s%d/%02d", home, ym[0], ym[1]))
	}
	return urls
}
