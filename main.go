package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

func createTopic(ctx context.Context, broker1Address, topic string, numPartitions, replicationFactor int) {
	conn, err := kafka.Dial("tcp", broker1Address)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to Kafka broker: %v", err))
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic(fmt.Sprintf("failed to get controller: %v", err))
	}

	controllerConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(fmt.Sprintf("failed to connect to controller: %v", err))
	}
	defer controllerConn.Close()

	topicConfigs := kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     numPartitions,
		ReplicationFactor: replicationFactor,
	}

	err = controllerConn.CreateTopics(topicConfigs)
	if err != nil {
		fmt.Printf("failed to create topic: %v\n", err)
	} else {
		fmt.Printf("topic %s created successfully\n", topic)
	}
}

func produce(ctx context.Context) {
	broker1Address := os.Getenv("KAFKA_BROKER_ADDRESS")
	topic := os.Getenv("KAFKA_TOPIC")

	if broker1Address == "" || topic == "" {
		panic("KAFKA_BROKER_ADDRESS or KAFKA_TOPIC is not set")
	}

	// initialize a counter
	i := 0

	// initialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker1Address},
		Topic:   topic,
	})

	for {
		// each kafka message has a key and value. The key is used
		// to decide which partition (and consequently, which broker)
		// the message gets published on
		err := w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte("this is message " + strconv.Itoa(i)), // create an arbitrary message payload for the value
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}

		// log a confirmation once the message is written
		fmt.Println("writes:", i)
		i++
		// sleep for a second
		time.Sleep(time.Second)
	}
}

func consume(ctx context.Context) {
	broker1Address := os.Getenv("KAFKA_BROKER_ADDRESS")
	topic := os.Getenv("KAFKA_TOPIC")

	if broker1Address == "" || topic == "" {
		panic("KAFKA_BROKER_ADDRESS or KAFKA_TOPIC is not set")
	}

	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker1Address},
		Topic:   topic,
		GroupID: "my-group",
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))
	}
}

func main() {
	// create a new context
	ctx := context.Background()

	broker1Address := os.Getenv("KAFKA_BROKER_ADDRESS")
	topic := os.Getenv("KAFKA_TOPIC")

	if broker1Address == "" || topic == "" {
		panic("KAFKA_BROKER_ADDRESS or KAFKA_TOPIC is not set")
	}

	// Create the topic
	createTopic(ctx, broker1Address, topic, 3, 1)

	// Produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	go produce(ctx)
	consume(ctx)
}
