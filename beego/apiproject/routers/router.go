// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"apiproject/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// Register routers.
	beego.Router("/", &controllers.MainController{})
	beego.Router("/shorten", &controllers.ShortController{})
	beego.Router("/expend", &controllers.ExpendController{})

}
