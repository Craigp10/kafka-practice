package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	fmt.Println("Running kafka")
	p := producer()
	defer p.Close()
	go consumer()

	fmt.Println("producing more, wait for 5 secords for consumer to receive inital messages")
	time.Sleep(5 * time.Second)
	i := 1
	topic := "myTopic"
	for {
		fmt.Println("Delivering message", i)
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(fmt.Sprintf("additional topic message %d", i)),
		}, nil)
		time.Sleep(2 * time.Second)
		i++
	}
}

func producer() *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	// No longer close within the producer() routine, the producer will be closed with the main routine closes.
	// defer p.Close()
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := "myTopic"
	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}

	// Wait for message deliveries before shutting down
	// p.Flush(15 * 1000)
	return p
}

func consumer() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"myTopic"}, nil)

	// A signal handler or similar could be used to set this to false to break the loop.
	run := true

	for run {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else if !strings.Contains(err.Error(), "Timed out") {
			// } else if !strings.Contains(err.(kafka.Error).Error(), "Timed out") {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
		// else {
		// 	// Needs further looking into to understand how to properly handle errors like timeout for kafka consumers.
		// 	fmt.Printf("time out: %v (%v)\n", err, msg)
		// 	//  e := kafka.NewError(err.(kafka.Error).Code(), err.(kafka.Error).String(), err.(kafka.Error).IsFatal())
		// 	run = false
		// }
	}

	c.Close()
}
