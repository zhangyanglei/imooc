package main

import "imooc/imooc-rabbitmq/RabbitMQ"

func main() {
	imoocOne := RabbitMQ.NewRabbitMQRouting("exImooc", "imooc_two")
	imoocOne.ReceiveRouting()
}
