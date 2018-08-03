// app
package controllers

import (
	"github.com/astaxie/beego"
)

type baseController struct {
	beego.Controller
}

type AppController struct {
	baseController
}

//处理get请求
func (this *AppController) Get() {
	this.TplName = "welcome.html"
}

// 处理post请求
func (this *AppController) Join() {

	uname := this.GetString("uname")
	tech := this.GetString("tech")

	// Check valid.
	if len(uname) == 0 {
		this.Redirect("/", 302)
		return
	}

	switch tech {
	case "longpolling":
		this.Redirect("/lp?uname="+uname, 302)
	case "websocket":
		this.Redirect("/ws?uname="+uname, 302)
	default:
		this.Redirect("/", 302)
	}

	// Usually put return after redirect.
	return
}
