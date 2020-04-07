package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strconv"
)

const MQURL = "amqp://leo:leo@localhost:5672/develop"

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	// queue Name
	QueueName string
	Exchange  string
	key       string `json:"key"`
	Mqurl     string
}

// 创建rabbitMQ结构体
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{QueueName: queueName, Exchange: exchange, key: key, Mqurl: MQURL}
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.FailOnError(err, "conn fail")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.FailOnError(err, "channel fail")
	return rabbitmq
}

// 断开链接
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

func (r *RabbitMQ) FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}

// 1. 创建简单模式下rabbitmq实例
// 每个rabbitmq结构体都会创建conn和channel, 直接初始化的时候做到这些就好了呀~
func NewRabbitmqSimple(queueName string) *RabbitMQ {
	return NewRabbitMQ(queueName, "", "")
}

// 2. 简单模式下publish信息
func (r *RabbitMQ) PublishSimple(message string, index int) {
	//1.申请队列
	_, err := r.channel.QueueDeclare(r.QueueName, false, false, false, false, nil)
	if err != nil {
		log.Printf("%s", err)
	}
	//2.发送消息到队列中
	r.channel.Publish(r.Exchange, r.QueueName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message + "  " + strconv.Itoa(index)),
	})
}

// 简单模式下receive消息
func (r *RabbitMQ) ConsumeSimple() {
	// 1.request queue
	_, err := r.channel.QueueDeclare(r.QueueName, false, false, false, false, nil)
	if err != nil {
		log.Printf("%s", err)
	}
	// 2.consume queue
	msgs, err := r.channel.Consume(r.QueueName, "", true, false, false, false, nil)
	if err != nil {
		log.Printf("%s", err)
	}
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Println(string(d.Body))
		}
	}()
	log.Println("等待接收消息...")
	<-forever // 负责让程序阻塞
	// 要让程序死锁, 所有的goroutine都要死锁, 匿名函数里面的goroutine没有死锁

}

// 订阅模式下创建结构体
func NewRabbitmqPubSub(exchangeName string) *RabbitMQ {
	return NewRabbitMQ("", exchangeName, "")
}

// 订阅模式生产
func (r *RabbitMQ) PublishSubscribe(message string, i int) {
	// 1.申请exchange
	err := r.channel.ExchangeDeclare(r.Exchange, "fanout", true, false, false, false, nil)
	r.FailOnError(err, "常见exchange错误")
	// 2.发送消息
	err = r.channel.Publish(r.Exchange, "fanout", false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(message + strconv.Itoa(i))})
	r.FailOnError(err, "exchange发送消息失败")
}

// 订阅模式消费
func (r *RabbitMQ) ConsumeSubscribe() {
	// 1.申请exchange
	err := r.channel.ExchangeDeclare(r.Exchange, "fanout", true, false, false, false, nil)
	r.FailOnError(err, "申请exchange发生错误")
	// 2.申请queue
	q, err := r.channel.QueueDeclare("", false, false, true, false, nil)
	r.FailOnError(err, "")
	// 3.绑定exchange和队列
	r.channel.QueueBind(q.Name, "", r.Exchange, false, nil)
	// 4. 开始消费
	msgs, err := r.channel.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Printf("%s", err)
	}
	c := make(chan bool)
	// 为什么没有死锁退出呢? 难道是因为<-chan有方向?
	go func() {
		for d := range msgs {
			fmt.Println(string(d.Body))
		}
	}()
	log.Println("等待接收消息...")
	<-c // 负责让程序阻塞
}

// 路由模式 type和key
func NewRabbitMQrouting(exchangeName string, routingKey string) *RabbitMQ {
	return NewRabbitMQ("", exchangeName, routingKey)
}

// 路由模式下发送消息
func (r *RabbitMQ) PublishRouting(message string) {
	// 1.尝试创建交换机
	err := r.channel.ExchangeDeclare(r.Exchange, "direct", true, false, false, false, nil)
	r.FailOnError(err, "Failed to declare a exchange")
	// 2.尝试发送消息
	err = r.channel.Publish(r.Exchange, r.key, false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(message)})
	r.FailOnError(err, "")
}

// 路由模式下消息的接收
func (r *RabbitMQ) ReceiveRouting() {
	// 1.尝试创建交换机
	err := r.channel.ExchangeDeclare(r.Exchange, "direct", true, false, false, false, nil)
	r.FailOnError(err, "")
	// 2.试探性创建队列
	q, err := r.channel.QueueDeclare("", false, false, true, false, nil)
	r.FailOnError(err, "")
	// 3.绑定队列和exchange
	err = r.channel.QueueBind(q.Name, r.key, r.Exchange, false, nil)
	// 开始消费
	msgs, err := r.channel.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Printf("%s", err)
	}
	c := make(chan bool)
	// 为什么没有死锁退出呢? 难道是因为<-chan有方向?
	go func() {
		for d := range msgs {
			fmt.Println(string(d.Body))
		}
	}()
	log.Println("等待接收消息...")
	<-c // 负责让程序阻塞

}
