// send_topic
//go run send_topic.go "kern.critical" "A critical kernal error"
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	conn, err := amqp.Dial("amqp:guest:guest@localhost:5672/")
	failOnError(err, "连接失败")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "创建通道失败")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_topic", //name
		"topic",      //type
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "创建交换器失败")
	body := bodyFrom(os.Args)
	err = ch.Publish(
		"logs_topic",          // exchange
		severityFrom(os.Args), //routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "发送消息失败")

}

//需两个以上的参数，第一个参数是路由键，第二个参数是发送的消息
func bodyFrom(args []string) string {
	var s string
	if (len(args) < 3) || os.Args[2] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[2:], " ")
	}

	return s
}

func severityFrom(args []string) string {
	var s string
	if len(args) < 2 || args[1] == "" {
		s = "info"
	} else {
		s = os.Args[1]
	}

	return s
}
