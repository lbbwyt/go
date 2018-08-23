// emit_log
//通过搭建一个日志系统来阐述发布/订阅模式，它包含两部分内容：一个用于产生日志消息的程序，另一个用于接收和打印消息。fanout交换器应用范例
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

//同一的错误处理
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "创建通道失败")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "创建交换器失败")

	body := bodyFrom(os.Args)

	err = ch.Publish(
		"logs", //交换器， 将消息发送到logs交换器
		"",     //路由键
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plan",
			Body:        []byte(body),
		},
	)
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		//默认发送hello
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
