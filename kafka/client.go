package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

var TOPIC = "order2"
var PARTITION = 2

func NewClient(topic string ,partition int) *kafka.Conn {
	conn,err := kafka.DialLeader(context.Background(),"tcp","localhost:9092",topic,partition)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
