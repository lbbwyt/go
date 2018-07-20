//如何通过表单处理文件上传
package webserver

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

//通过表单处理文件上传
func Upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		//get请求直接跳转到用户名密码输入页面
		//以当前时间生成一个唯一的token值
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		//使用r.FormFile获取文件句柄，然后对文件进行存储等处理。
		t, _ := template.ParseFiles("html/upload.gtpl")
		t.Execute(w, token)
	} else {
		//		参数表示maxMemory，调用ParseMultipartForm之后，上传的文件存储在maxMemory大小的内存里面，
		//		如果文件大小超过了maxMemory，那么剩下的部分将存储在系统的临时文件中。
		r.ParseMultipartForm(32 << 20)
		//		使用r.FormFile获取文件句柄，然后对文件进行存储等处理。
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		//fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./img/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		fmt.Fprintf(w, "上传成功")
	}
}
