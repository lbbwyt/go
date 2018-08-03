// short
package controllers

import (
	"apiproject/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
)

var (
	urlcache cache.Cache
)

func init() {
	urlcache, _ = cache.NewCache("memory", `{"interval":0}`)
}

type ShortResult struct {
	UrlShort string
	UrlLong  string
}

type ShortController struct {
	beego.Controller
}

//长转短
func (this *ShortController) Get() {
	var result ShortResult
	longurl := this.Input().Get("longurl")
	beego.Info(longurl)
	result.UrlLong = longurl
	urlmd5 := models.GetMd5(longurl)
	beego.Info(urlmd5)
	if urlcache.IsExist(urlmd5) {
		result.UrlShort = urlcache.Get(urlmd5).(string)
	} else {
		result.UrlShort = models.Generate()
		err := urlcache.Put(urlmd5, result.UrlShort, 0)
		if err != nil {
			beego.Info(err)
		}
		//用于短转长直接从缓冲里面取
		err = urlcache.Put(result.UrlShort, longurl, 0)
		if err != nil {
			beego.Info(err)
		}
	}

	this.Data["json"] = result
	this.ServeJSON()
}
