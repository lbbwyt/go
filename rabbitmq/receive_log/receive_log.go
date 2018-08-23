// receive_log
//如果想将日志消息保存到文件，只需在命令终端中执行下面的命令：
//go run receive_logs.go > logs_from_rabbit.log
package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "连接失败")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "创建通道失败")
	defer ch.Close()
	//交换器的定义需和消息的发送方一致
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
	//声明一个临时队列，
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "创建交换器失败")

	//将临时队列绑定到logs交换器上
	err = ch.QueueBind(
		q.Name, //队列名称
		"",     //路由键
		"logs", //交换器
		false,
		nil,
	)
	failOnError(err, "绑定临时队列带logs交换器失败")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "注册消费者失败")
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

}
