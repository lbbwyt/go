// webserver_test.go
package webserver

import (
	"log"
	"net/http"
	"testing"
)

//url为http://localhost:9090/?url_long=111&url_long=222
func TestSayHelloName(t *testing.T) {
	http.HandleFunc("/shutdown", ShutDown)
	http.HandleFunc("/", SayHelloName)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/upload", Upload)       //设置访问的路由
	err := http.ListenAndServe(":9091", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
