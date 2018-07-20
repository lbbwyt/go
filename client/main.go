// client project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
)

func main() {
	//os.Args
	//程序获取运行他时给出的参数，如： go run main.go :6700
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage : %s host:port", os.Args[0])
		os.Exit(1)
	}
	//首先程序将用户的输入作为参数即service传入net.ResolveTCPAddr获取一个tcpAddr,（tcpAddr封装了ip和端口）
	//然后把tcpAddr传入DialTCP后创建了一个TCP连接conn，通过conn来发送请求信息，
	//最后通过ioutil.ReadAll从conn中读取全部的文本，也就是服务端响应反馈的信息
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkErr(err)
	fmt.Println("tcpAddr的ip为：" + tcpAddr.IP.String())
	fmt.Println("tcpAddr的port为：" + strconv.Itoa(tcpAddr.Port))
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkErr(err)
	defer conn.Close()
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkErr(err)
	result, err := ioutil.ReadAll(conn)
	checkErr(err)
	fmt.Println("服务端返回：" + string(result))
	os.Exit(0)
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
