// spider
package spider

//这里假设接口GetBody直接可以抓取"https://golang.org"首页
//的“Build version”字段来得到当前Golang发布出来的版本。
type Spider interface {
	GetBody() string
}
