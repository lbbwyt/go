// expend
package controllers

import (
	"github.com/astaxie/beego"
)

type ExpendController struct {
	beego.Controller
}

//短转长（直接从缓冲里取）05ftgH
func (this *ExpendController) Get() {
	var result ShortResult
	shorturl := this.Input().Get("shorturl")
	beego.Info(shorturl)
	result.UrlShort = shorturl
	if urlcache.IsExist(shorturl) {
		result.UrlLong = urlcache.Get(shorturl).(string)
	} else {
		result.UrlLong = ""
	}
	this.Data["json"] = result
	this.ServeJSON()
}
