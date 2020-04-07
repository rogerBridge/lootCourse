package main

import (
	"example/RabbitMQ"
	"strconv"
)

func main() {
	//rabbitmq := RabbitMQ.NewRabbitmqSimple("leo")
	////done := make(chan bool)
	////wg := new(sync.WaitGroup)
	//t0 := time.Now()
	////for i := 0; i < 10000; i++ {
	////	go rabbitmq.PublishSimple("Hello Index :)", i, done)
	////	<-done
	////}
	//for i := 0; i < 1000; i++ {
	//	//wg.Add(1)
	//	rabbitmq.PublishSimple("Hello Index :)", i)
	//}
	////wg.Wait()
	//fmt.Println("所有线程跑完咯~")
	////for v := range done{
	////	fmt.Println(v)
	////}
	//fmt.Println("time consume is", time.Since(t0))

	// 订阅模式发送
	//rabbitmq := RabbitMQ.NewRabbitmqPubSub("mox")
	//for i:=0; i<100; i++ {
	//	rabbitmq.PublishSubscribe("发送的消息的序列号是:", i)
	//}

	// 路由模式发送
	key1 := RabbitMQ.NewRabbitMQrouting("ex", "Leborn")
	key2 := RabbitMQ.NewRabbitMQrouting("ex", "Leonard")
	for i := 0; i < 10; i++ {
		key1.PublishRouting("routing key: Leborn " + strconv.Itoa(i))
		key2.PublishRouting("routing key: Leonard " + strconv.Itoa(i))
	}
}
