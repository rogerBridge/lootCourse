package main

import "example/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitmqSimple("leo")
	rabbitmq.ConsumeSimple()
}
