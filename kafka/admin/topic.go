package admin

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

func Exec()  {
	conn,err := kafka.Dial("tcp","localhost:9092")
	if err != nil {
		log.Fatal(err)
	}

	brokers,err := conn.Brokers()
	if err != nil {
		log.Fatal(err)
	}

	for _,v := range brokers{
		fmt.Println(v.ID,v.Host,v.Port)
	}
}