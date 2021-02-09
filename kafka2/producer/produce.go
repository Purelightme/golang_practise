package producer

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)

var topic string  = "order-sarama-admin"

func Exec()  {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner
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

func AsyncProduce()  {
	config := sarama.NewConfig()
	client,err := sarama.NewClient([]string{"localhost:9092","localhost:9093"},config)
	if err != nil {
		log.Fatal(err)
	}

	producer,err := sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

	msg := "{\"id\":2,\"type\":\"async-sarama\"}"
	producer.Input() <- &sarama.ProducerMessage{Topic:topic,Key:sarama.StringEncoder("asyncKey"),Value:sarama.StringEncoder(msg)}

	select {
	case err = <- producer.Errors():
		fmt.Println(err)
	case ret := <- producer.Successes():
		fmt.Println("发送成功,",ret)
	default:
		fmt.Println("-----")
	}
}