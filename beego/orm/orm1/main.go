// orm1 project main.go
package main

import (
	"fmt"
	"orm1/models"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbstring = "root:123456@tcp(127.0.0.1:3306)/world?charset=utf8"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dbstring, 30)
	//	创建表
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	profile := new(models.Profile)
	profile.Age = 30

	user := new(models.User)
	user.Profile = profile
	user.Name = "slene"

	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))

	var users []*models.User
	qs := o.QueryTable("user")
	num, err := qs.Filter("name__contains", "sl").All(&users)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)
	if num > 0 {
		for _, item := range users {
			fmt.Println(item.Name)
		}
	}

	//	构造查询
	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Select("user.name",
		"profile.age").
		From("user").
		InnerJoin("profile").On("user.profile_id = profile.id").
		Where("age > ?").
		OrderBy("name").Desc().
		Limit(10).Offset(0)

	// 导出 SQL 语句
	sql := qb.String()
	fmt.Println("*************")
	fmt.Println(sql)
	// 执行 SQL 语句
	num, err = o.Raw(sql, 20).QueryRows(&users)
	fmt.Println(num)

	//自定义条件表达式
	//	cond := orm.NewCondition()
	//	cond1 := cond.And("profile__isnull", false).AndNot("status__in", 1).Or("profile__age__gt", 2000)

	//	qs := orm.QueryTable("user")
	//	qs = qs.SetCond(cond1)
	//// WHERE ... AND ... AND NOT ... OR ...

	//	cond2 := cond.AndCond(cond1).OrCond(cond.And("name", "slene"))
	//	qs = qs.SetCond(cond2).Count()
	//// WHERE (... AND ... AND NOT ... OR ...) OR ( ... )

}
