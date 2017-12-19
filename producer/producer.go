package main

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

var (
	addrs = []string{"localhost:9092"}
	topic = "test"
)

func main() {
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(addrs, conf)
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

	for index := 0; index < 1000; index++ {
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(fmt.Sprintf("message: %d", index)),
		}

		if _, _, err := producer.SendMessage(msg); err != nil {
			fmt.Println(err)
		}
	}
}
