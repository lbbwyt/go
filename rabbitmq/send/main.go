// send project main.go
//消息的生产者
package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

//在Go语言中经常需要使用if语句来检查操作结果，为了避免在代码中到处散落if(err != nil)语句，可以使用下列方法
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "连接rabbitmq服务器失败")
	defer conn.Close()
	//在连接上创建通道，之后我们的大多数API操作都是围绕通道来实现的：
	ch, err := conn.Channel()
	failOnError(err, "打开通道失败")
	defer ch.Close()

	//定义一个队列用来存储、转发消息，然后我们的sender只需要将消息发送到这个队列中，就完成了消息的publish操作
	q, err := ch.QueueDeclare(
		"hello", //name
		true,    //durable， durable参数在生产者和消费者程序中都要指定为True
		false,   //delete when unused
		false,   //exclusive
		false,   //no wait
		nil,     //arguments
	)
	failOnError(err, "创建队列失败")
	body := "this is a msg send to hello queue by default exchange"
	err = ch.Publish(
		"",     //交换器 默认的交换器,消息将会根据routing_key所指定的参数(这里就是队列的名称)进行查找，如果存在就会分发到相应的队列。
		q.Name, //路由键
		false,
		false,
		amqp.Publishing{
			//消息持久化
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plan",
			Body:         []byte(body),
		},
	)
	failOnError(err, "发送消息失败")

}
