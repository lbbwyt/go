package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

const url = "https://localhost:8000"

//通过命令行接受协议版本号
var httpVersion = flag.Int("version", 2, "HTTP version")

func main() {
	flag.Parse()
	client := &http.Client{}
	//    server.crt为服务端证书
	caCert, err := ioutil.ReadFile("server.crt")
	if err != nil {
		log.Fatalf("读取服务端证书失败: %s", err)
	}
	//创建一个证书池,
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	// 创建tls配置项
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	switch *httpVersion {
	case 1:
		//Transport可用于缓冲连接，以便充分使用
		//Transport并发安全的
		client.Transport = &http.Transport{
			TLSClientConfig: tlsConfig,
		}
	case 2:
		client.Transport = &http2.Transport{
			TLSClientConfig: tlsConfig,
		}
	}

	resp, err := client.Get(url)

	if err != nil {
		log.Fatalf("%s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("读取返回信息失败%s", err)
	}
	fmt.Printf("获取返回信息 %d: %s %s\n", resp.StatusCode, resp.Proto, string(body))
}
