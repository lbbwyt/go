//使用Go搭建一个Web服务器，并实现简单的登录操作，
package webserver

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func SayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello world!") //这个写入到w的是输出到客户端的
}

//BUG(libaobao) 验证逻辑未实现
func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		//get请求直接跳转到用户名密码输入页面
		//以当前时间生成一个唯一的token值
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("html/login.gtpl")
		t.Execute(w, token)
	} else {

		token := r.Form.Get("token")
		if token != "" {
			//验证token的合法性, 多次提交时这里的token的值时一样的，因此可以用来判断是否为重复提交
			fmt.Println(token)
		} else {
			//不存在token报错
		}
		//post请求，那么执行登陆的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		//暂停一秒，以便有足够时间测试多次提交
		time.Sleep(time.Second * 1)
		fmt.Fprintf(w, "登录成功")
	}
}

func ShutDown(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Goodbye!")
	//退出后http服务的端口就不会被占用
	os.Exit(0)
}
