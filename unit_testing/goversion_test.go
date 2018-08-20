// goversion_test
package main

import (
	"testing"
	"unit_testing/spider"

	"github.com/golang/mock/gomock"
)

//mockgen工具是gomock提供的用来为要mock的接口生成实现的
//-source： 指定接口文件
//-destination: 生成的文件名
//-package:生成文件的包名
//-imports: 依赖的需要import的包
//-aux_files:接口文件不止一个文件时附加文件
//-build_flags: 传递给build工具的参数

//mockgen -destination spider/mock_spider.go -package spider -source spider/spider.go
func TestGetGoVersion(t *testing.T) {
	//这里spider.CreateGoVersionSpider()返回一个实现了Spider接口的用来获得Go版本号的爬虫。
	//这个单元测试其实既测试了函数GetGoVersion也测试了
	//spider.CreateGoVersionSpider返回的对象。
	//而有时候，我们可能仅仅想测试下GetGoVersion函数，
	//或者我们的spider.CreateGoVersionSpider爬虫实现还没有写好，那该如何是好呢？
	//	此时Mock工具就显的尤为重要了。

	//	v := GetGoVersion(spider.CreateGoVersionSpider())

	//****************这里是分割下***************
	//	t.Skip("跳过基准测试后下面的测试将不再执行")
	//	fmt.Println("跳过基准测试")

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	//NewMockSpider返回一个接口的mock实例(该方法是mockgen工具自动生成的)
	mockSpider := spider.NewMockSpider(mockCtl)
	//	这里EXPECT()得到实现的对象，然后调用实现对象的接口方法
	mockSpider.EXPECT().GetBody().Return("go1.8.3")
	v := GetGoVersion(mockSpider)

	if v != "go1.8.3" {
		t.Error("Get wrong version %s", v)
	}
}
