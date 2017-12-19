package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
	"github.com/heppu/kafka-demo/consumer/db"
)

var (
	addrs = []string{"localhost:9092"}
	topic = "test"
)

func main() {
	consumer, err := sarama.NewConsumer(addrs, sarama.NewConfig())
	if err != nil {
		log.Fatal("Could not create kafka consumer:", err)
	}
	defer consumer.Close()

	partCosumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal("Could not start consuming partition:", err)
	}
	defer partCosumer.Close()

	log.Println("Kafka consumer listening topic", topic)

	// Hook into SIGINT to shutdown server gracefully
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Label that allows us to break from select and for loop both
ConsumerLoop:
	for {
		select {
		case msg := <-partCosumer.Messages():
			log.Printf("Received message: %s\n", msg.Value)
			go saveMessage(msg)
		case <-signals:
			log.Println("Received SIGINT shutting down gracefully")
			break ConsumerLoop // Break from outter for loop
		}
	}
}

func saveMessage(kafkaMsg *sarama.ConsumerMessage) {
	msg := &db.Message{Message: string(kafkaMsg.Value)}
	if err := db.Client.InsertMessage(msg); err != nil {
		log.Println("Could not insert message to DB:", err)
		return
	}
	log.Printf("Message writen to db with ID: %d", msg.ID)
}
