// receive_topic
//go run receive_topic.go "*.critical"
package main

import (
	"fmt"
	"log"
	"os"

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
	//通道和发送方声明的通道一致
	err = ch.ExchangeDeclare(
		"logs_topic", //name
		"topic",      //type
		true,
		false,
		false,
		false,
		nil,
	)
	//	声明临时队列
	q, err := ch.QueueDeclare(
		"", //name
		false,
		false,
		true,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	if len(os.Args) < 2 {
		log.Printf("Usage: %s [info] [warning] [error]", os.Args[0])
		os.Exit(0)
	}
	//路由键通过参数传递
	for _, s := range os.Args[1:] {
		log.Printf("Binding queue %s to exchange %s with routing key %s",
			q.Name, "logs_topic", s)
		err = ch.QueueBind(
			q.Name,       //队列名称
			s,            //路由键
			"logs_topic", //交换器
			false,
			nil)
		failOnError(err, "Failed to bind a queue")
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

}
