// models
package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id      int
	Name    string
	Profile *Profile `orm:"rel(one)"`      // OneToOne relation
	Post    []*Post  `orm:"reverse(many)"` // 设置一对多的反向关系
}

//// 多字段索引
//func (u *User) TableIndex() [][]string {
//	return [][]string{
//		[]string{"Id", "Name"},
//	}
//}

//// 多字段唯一键
//func (u *User) TableUnique() [][]string {
//	return [][]string{
//		[]string{"Name", "Email"},
//	}
//}

//扼要描述
type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"` //设置一对多关系
	Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`

	//	Created time.Time `orm:"auto_now_add;type(date)"`
	//  Created time.Time `orm:"auto_now_add;type(datetime)"`
	//	Money float64 `orm:"digits(12);decimals(4)"`
	//	Name string `orm:"column(user_name)"`

}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User), new(Post), new(Profile), new(Tag))
}
