// orm project main.go
package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

const (
	dbstring = "root:123456@tcp(127.0.0.1:3306)/world?charset=utf8"
)

func init() {
	orm.RegisterDataBase("default", "mysql", dbstring, 30)
	//	注册模型
	orm.RegisterModel(new(User))
	//	创建表
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()
	user := User{Name: "lbbwyt"}

	//	插入
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	//更改
	user.Name = "zhangsan"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	//查询
	u := User{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("name : v%, ERR: %v\n", u.Name, err)

	//	//删除
	//	num, err = o.Delete(&u)
	//	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	//使用原生的sql查询
	var maps []orm.Params
	num, err = o.Raw("SELECT * FROM user").Values(&maps)
	for _, item := range maps {
		fmt.Println(item["id"], ":", item["name"])
	}

}
