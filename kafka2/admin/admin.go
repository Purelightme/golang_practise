package admin

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)

func Exec(){
	topic := "order-sarama-admin"
	config := sarama.NewConfig()
	admin,err := sarama.NewClusterAdmin([]string{"localhost:9092","localhost:9093"},config)
	if err != nil {
		log.Fatal(err)
	}
	defer admin.Close()

	err = admin.CreateTopic(topic,&sarama.TopicDetail{
		NumPartitions:4,
		ReplicationFactor:2,
	},false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("创建topic成功,topic:",topic)
}

func ListTopics()  {
	config := sarama.NewConfig()
	admin,err := sarama.NewClusterAdmin([]string{"localhost:9092","localhost:9093"},config)
	if err != nil {
		log.Fatal(err)
	}
	defer admin.Close()

	topics,err := admin.ListTopics()
	if err != nil {
		log.Fatal(err)
	}
	for name,detail := range topics{
		res,err := json.Marshal(detail)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(name,string(res))
	}
}
