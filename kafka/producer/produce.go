package producer

import (
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	myKafka "purelight/kafka"
	"time"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func Exec()  {
	var err error
	var bytes  int
	var user = User{Id:1,Name:"purelightme"}

	conn := myKafka.NewClient(myKafka.TOPIC,myKafka.PARTITION)
	defer conn.Close()

	err = conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	if err != nil {
		log.Fatal(err)
	}

	err = conn.SetRequiredAcks(1)
	if err != nil {
		log.Fatal(err)
	}

	order,err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	bytes,err = conn.WriteMessages(
		kafka.Message{Value: []byte("one")},
		kafka.Message{Value: []byte("two")},
		kafka.Message{Value: []byte("three")},
		kafka.Message{Value: []byte(order)},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("成功生产",bytes,"个字节消息")
}
