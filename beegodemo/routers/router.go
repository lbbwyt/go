package routers

import (
	"beegodemo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// Register routers.
	beego.Router("/", &controllers.AppController{})
	// 自定义方法名，对于post请求将执行Join方法
	//未指定方法名时，get请求将执行Get方法
	beego.Router("/join", &controllers.AppController{}, "post:Join")

	// Long polling.
	beego.Router("/lp", &controllers.LongPollingController{}, "get:Join")
	beego.Router("/lp/post", &controllers.LongPollingController{})
	beego.Router("/lp/fetch", &controllers.LongPollingController{}, "get:Fetch")

	// WebSocket.
	beego.Router("/ws", &controllers.WebSocketController{})
	beego.Router("/ws/join", &controllers.WebSocketController{}, "get:Join")
}
