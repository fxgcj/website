package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/ckeyer/commons/lib"
	"net/http"
)

const (
	KEY_A = "a"
	KEY_B = "b"

	COOKIE_IS_LOGINED = "is_logined"
)

type AdminBaseController struct {
	BaseController
}

func (a *AdminBaseController) Prepare() {
	if !a.Auth() {
		if a.Ctx.Input.Url() != "/admin/login" {
			a.Redirect("/admin/login", http.StatusMovedPermanently)
			return
		}
	} else {
		if a.Ctx.Input.Url() == "/admin/login" {
			a.Redirect("/admin", http.StatusMovedPermanently)
			return
		}
	}
	a.InitWebPage()
}

func (a *AdminBaseController) Auth() bool {
	if value, ok := a.GetCookie(COOKIE_IS_LOGINED); ok && value == "true" {
		return true
	}
	return false
}

func (a *AdminBaseController) setA_B() (key_a, key_b int) {
	key_a = lib.RandomInt(5, 49)
	a.SetSession(KEY_A, key_a)
	a.Data["a"] = key_a
	key_b = lib.RandomInt(5, 50)
	a.SetSession(KEY_B, key_b)
	a.Data["b"] = key_b
	return
}

func (a *AdminBaseController) verifyA_B(sec_hash string) bool {
	defer func() {
		if r := recover(); r != nil {
			log.Error(r)
		}
	}()
	key_a := a.GetSession(KEY_A).(int)
	key_b := a.GetSession(KEY_B).(int)

	if commitSec(website.CommitPassword, key_a, key_b) == sec_hash {
		return true
	}
	log.Error("auth error")
	log.Debug("get session a=", key_a)
	log.Debug("get session b=", key_b)
	return false
}

func (a *AdminBaseController) verifySecret() {
	m := make(map[string]int)
	m["a"] = lib.RandomInt(5, 49)
	a.SetSession("a", m["a"])
	m["b"] = lib.RandomInt(5, 50)
	a.SetSession("b", m["b"])

	a.WriteJSON(http.StatusOK, m)
}

func commitSec(sec string, a, b int) string {
	h := md5.New()
	h.Write([]byte(fmt.Sprint(sec, a+b)))
	return hex.EncodeToString(h.Sum(nil))
}
