package main

import (
	"example/RabbitMQ"
	"fmt"
	"time"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitmqSimple("leo")
	//done := make(chan bool)
	//wg := new(sync.WaitGroup)
	t0 := time.Now()
	//for i := 0; i < 10000; i++ {
	//	go rabbitmq.PublishSimple("Hello Index :)", i, done)
	//	<-done
	//}
	for i := 0; i < 10000; i++ {
		//wg.Add(1)
		rabbitmq.PublishSimple("Hello Index :)", i)
	}
	//wg.Wait()
	fmt.Println("所有线程跑完咯~")
	//for v := range done{
	//	fmt.Println(v)
	//}
	fmt.Println("time consume is", time.Since(t0))
	// 开始并发发送10000条消息
}
