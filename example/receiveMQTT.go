package main

import "example/RabbitMQ"

func main() {
	//rabbitmq := RabbitMQ.NewRabbitmqSimple("leo")
	//rabbitmq.ConsumeSimple()

	//rabbitmq := RabbitMQ.NewRabbitmqPubSub("mox")
	//rabbitmq.ConsumeSubscribe()

	// 路由模式接收
	key1 := RabbitMQ.NewRabbitMQrouting("ex", "Leborn")
	key1.ReceiveRouting()
}
