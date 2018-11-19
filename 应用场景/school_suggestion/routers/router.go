package routers

import (
	"school_suggestion/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/shcool", &controllers.SchoolController{})
	beego.Router("/shcool/SerchSchool", &controllers.SchoolController{}, "*:SerchSchool")
}
