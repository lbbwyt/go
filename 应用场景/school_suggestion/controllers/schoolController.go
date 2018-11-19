package controllers

import (
	"fmt"
	"school_suggestion/logic"
	"time"

	"github.com/astaxie/beego"
)

type SchoolController struct {
	beego.Controller
}

func (c *SchoolController) Get() {
	c.TplName = "school.tpl"
}

func (c *SchoolController) SerchSchool() {
	uname := c.GetString("name")
	start := time.Now().UnixNano()
	//schools := logic.SearchV2(keyword, 16)
	schools := logic.Search(uname, 16)
	end := time.Now().UnixNano()
	fmt.Printf("keyword:%s result:%d cost:%d us\n", uname, len(schools), (end-start)/1000)
	c.Data["json"] = schools
	c.ServeJSON()
}
