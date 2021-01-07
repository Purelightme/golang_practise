package consumer

import (
	context2 "context"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var wg sync.WaitGroup

type Handler struct {
	
}

func (c *Handler)Setup(session sarama.ConsumerGroupSession) error  {
	return nil
}

func (c *Handler)Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Handler)ConsumeClaim(session sarama.ConsumerGroupSession,claim sarama.ConsumerGroupClaim) error  {

	for messages := range claim.Messages() {
		log.Println(messages.Topic,messages.Partition,messages.Offset,string(messages.Key),string(messages.Value),
			messages.Timestamp.Format("2006-01-02 15:04:05"))
		session.MarkMessage(messages,"")
	}

	return nil
}


func Exec() {
	config := sarama.NewConfig()
	consume, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatal(err)
	}
	defer consume.Close()

	cp, err := consume.ConsumePartition("order", 0, 15)
	if err != nil {
		log.Fatal(err)
	}
	defer cp.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0

consumeLoop:
	for {
		select {
		case msg := <-cp.Messages():
			fmt.Println(msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value), msg.Timestamp.Format("2006-01-02 15:04:05"))
			consumed++
		case <-signals:
			break consumeLoop
		}
	}
	fmt.Println("Consumed: ", consumed)
}

func GroupExec()  {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	consume, err := sarama.NewConsumerGroup([]string{"localhost:9092","localhost:9093"},"group1", config)
	if err != nil {
		log.Fatal(err)
	}
	defer consume.Close()
	context,cancel := context2.WithCancel(context2.Background())
	handler := Handler{}

	wg.Add(1)

	go func() {
		wg.Done()
		for {
			err := consume.Consume(context,[]string{"order-sarama-admin"},&handler)
			if err != nil {
				log.Fatal(err)
			}
			if context.Err() != nil {
				return
			}
		}
	}()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-context.Done():
		log.Println("terminating: context cancelled")
	case <-sigterm:
		log.Println("terminating: via signal")
	}
	cancel()
	wg.Wait()
}
