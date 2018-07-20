// server project main.go
package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	service := ":6700"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkErr(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkErr(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	//由于我们需要保持与客户端的长连接，所以不能在读取完一次请求后就关闭连接,只有在超时后才关闭连接
	//设置超时时间
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	// request在创建时需要指定一个最大长度以防止flood attack；
	request := make([]byte, 128)
	defer conn.Close()
	//	由于conn.SetReadDeadline()设置了超时，当一定时间内客户端无请求发送，conn便会自动关闭，下面的for循环即会因为连接已关闭而跳出。
	//	设置客户端是否和服务器端保持长连接，可以降低建立TCP连接时的握手开销，对于一些需要频繁交换数据的应用场景比较适用。
	for {
		read_len, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("客户端请求的为" + string(request))
		if read_len == 0 {
			break
		} else if string(request) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}
		//每次读取到请求处理完毕后，需要清理request，因为conn.Read()会将新读取到的内容append到原内容之后。
		request = make([]byte, 128)
	}

}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
