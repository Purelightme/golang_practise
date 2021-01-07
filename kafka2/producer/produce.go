package producer

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)

func Exec()  {
	topic := "order-sarama-admin"
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	client,err := sarama.NewClient([]string{"localhost:9092","localhost：9093"},config)
	if err != nil {
		log.Fatal(err)
	}

	producer,err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

	msg := "{\"id\":1,\"type\":\"sarama\"}"
	partition, offset, err := producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Key:sarama.StringEncoder("testKey"),
		Value:sarama.StringEncoder(msg),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("成功生产一条消息，topic:",topic,",partition:",partition,",offset:",offset)
}