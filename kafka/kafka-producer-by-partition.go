package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func main() {

	partitionSize := 2
	topicName := "topic_deneme"

	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	var messages = [7]string{"one", "two", "three", "for", "five", "six", "seven"}
	for i := 0; i < len(messages); i++ {

		partition := i % partitionSize
		conn, err = kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topicName, partition)
		if err != nil {
			log.Fatal("failed to dial leader:", err)
		}

		conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		_, err = conn.WriteMessages(
			kafka.Message{Value: []byte(messages[i])},
		)
		if err != nil {
			log.Fatal("failed to write messages:", err)
		}

		if err := conn.Close(); err != nil {
			log.Fatal("failed to close writer:", err)
		}

	}

}
