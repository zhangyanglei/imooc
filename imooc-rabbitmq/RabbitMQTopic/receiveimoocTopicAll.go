package main

import "imooc/imooc-rabbitmq/RabbitMQ"

func main() {
	imoocOne := RabbitMQ.NewRabbitMQTopic("exImoocTopic", "#")
	imoocOne.ReceiveTopic()
}
