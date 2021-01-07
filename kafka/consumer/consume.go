package consumer

import (
	"fmt"
	"log"
	myKafka "purelight/kafka"
	"time"
)
func Exec()  {
	var err error
	var length int
	conn := myKafka.NewClient(myKafka.TOPIC,myKafka.PARTITION)
	defer conn.Close()

	err = conn.SetReadDeadline(time.Now().Add(time.Second * 10))
	if err != nil {
		log.Fatal(err)
	}

	batch := conn.ReadBatch(10e3,1e6)
	defer batch.Close()

	for {
		b := make([]byte,10e3)
		length,err = batch.Read(b)
		if err != nil {
			break
		}

		fmt.Println(string(b[:length]))
	}

}