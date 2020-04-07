package main

import "example/RabbitMQ"

func main() {
	//rabbitmq := RabbitMQ.NewRabbitmqPubSub("mox")
	//rabbitmq.ConsumeSubscribe()

	// 路由模式接收
	key1 := RabbitMQ.NewRabbitMQrouting("ex", "Leonard")
	key1.ReceiveRouting()
}
