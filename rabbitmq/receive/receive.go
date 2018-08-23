// receive
//消息的接受者
package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "连接服务器失败")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "创建通道失败")
	defer ch.Close()
	//和发送方声明的queue一致
	q, err := ch.QueueDeclare(
		"hello", // name
		true,    // durable， durable参数在生产者和消费者程序中都要指定为True，
		//		RabbitMQ不允许创建多个名称相同而参数不同的队列，
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "创建队列失败")

	//假设有这么一个场景，存在两个消费者程序，所有的单数序列消息都是长耗时任务而双数序列消息则都是简单任务
	//	，那么结果将是一个消费者一直处于繁忙状态而另外一个则几乎没有任务被挂起

	//为了避免这种情况，我们可以给队列设置预取数(prefect count)为1。它告诉RabbitMQ不要一次性分发超过1个的消息给某一个消费者，换句话说，就是当分发给该消费者的前一个消息还没有收到ack确认时，
	//RabbitMQ将不会再给它派发消息，而是寻找下一个空闲的消费者目标进行分发,以实现公平分发。
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "传输队列中的消息失败")
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			//手动发送一个确认消息，也可以我们将Consume()函数的aotu-ack参数设为true
			//忘记对消息进行确认是一个比较常见的错误，后果很严重
			d.Ack(false)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	//从一个空的channel读取消息时，将会阻塞
	<-forever
}
