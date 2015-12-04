package conf

import (
	"io/ioutil"
	"os"

	"encoding/json"

	"sync"

	logpkg "github.com/fxgcj/website/lib/log"
)

var (
	confpath = "conf/v2.json"
	log      = logpkg.GetLogger()
	config   *Config
)

type Config struct {
	sync.RWMutex
	AppConfig *AppConfig `json:"app"`
	WebSite   *WebSite   `json:"website"`
	MongoDB   *MongoDB   `json:"mongodb"`
}

type AppConfig struct {
	RunMode    string `json:"mode"`
	Name       string `json:"name"`
	ServerName string `json:"server_name"`
	Port       int    `json:"port"`
}

type WebSite struct {
	Title          string            `json:"title"`
	Keywords       []string          `json:"keywords"`
	Description    string            `json:"description"`
	CommitPassword string            `json:"commit_password"`
	HostUrl        string            `json:"host_url"`
	FileUrl        string            `json:"file_url"`
	JsUrl          string            `json:"js_url"`
	CssUrl         string            `json:"css_url"`
	ImgUrl         string            `json:"img_url"`
	CustomJsUrl    string            `json:"custom_js_url"`
	CustomCssUrl   string            `json:"custom_css_url"`
	CustomImgUrl   string            `json:"custom_img_url"`
	EnableDomain   []string          `json:"enable_domain"`
	FriendLinks    map[string]string `json:"friend_links"`
}

type MongoDB struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

// LoadConf 加载配置文件
func LoadConf(path string) (c *Config, err error) {
	confpath = path
	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		log.Errorf("打开文件 %s 失败 %s...", path, err)
		return
	}

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Errorf("读取文件 %s 失败，%s...", path, err)
		return
	}
	c = new(Config)
	err = json.Unmarshal(bs, c)
	if err != nil {
		log.Errorf("解析文件 %s 失败, %s", path, err)
		return
	}
	return
}

// GetConf ...
func GetConf() *Config {
	if config != nil {
		return config
	}
	if confpath != "" {
		if c, err := LoadConf(confpath); err == nil {
			if c != nil {
				config = c
			}
			log.Noticef("加载配置 %s 成功", confpath)
			return config
		}
	} else {
		panic("path is nil ")
	}
	log.Errorf("%s, %#v", confpath, config)
	panic("获取配置异常...")
}
