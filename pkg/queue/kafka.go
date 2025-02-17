package queue

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/segmentio/kafka-go"
)

func KafkaConn(cfg *config.Config, topic string) *kafka.Conn {

	conn, err := kafka.DialLeader(context.Background(), "tcp", cfg.Kafka.Url, topic, 0)
	if err != nil {
		panic(err.Error())
	}
	return conn

}

func PushMessageToQueue(cfg *config.Config, key, topic string, message ...kafka.Message) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", cfg.Kafka.Url, topic, 0)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	for _, msg := range message {
		msg.Key = []byte(key)
		if _, err := conn.WriteMessages(msg); err != nil {
			log.Printf("Error writing message: %s", err.Error())
		}
	}

}

func KafkaReader(topic, groupID string) *kafka.Reader {

	readerConfig := kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		GroupID:     groupID,
		Topic:       topic,
		StartOffset: kafka.LastOffset,
	}

	reader := kafka.NewReader(readerConfig)

	fmt.Println("Starting consumer...")

	return reader
}

// func ReadMessages(conn *kafka.Conn, key string) {
// 	for {

// 		msg, err := conn.ReadMessage(10e6) // 10e6 is the maximum size of the message to read
// 		if err != nil {
// 			fmt.Println("Error reading message:", err)
// 			break
// 		}

// 		if string(msg.Key) == key {
// 			fmt.Printf("Received message: key=%s, value=%s\n", string(msg.Key), string(msg.Value))
// 		}
// 	}
// }

// func ReadMessages(reader *kafka.Reader, Key string) {
// 	for {
// 		msg, err := reader.ReadMessage(context.Background())
// 		if err != nil {
// 			fmt.Println("Error reading message:", err)
// 			break
// 		}

// 		if string(msg.Key) == Key {
// 			fmt.Printf("Received message: key=%s, value=%s\n", string(msg.Key), string(msg.Value))
// 			// Process the message
// 		} else {
// 			fmt.Printf("Skipping message with key=%s\n", string(msg.Key))
// 		}
// 	}
// }

// func IsTopicIsAlreadyExits(conn *kafka.Conn, topic string) bool {
// 	partition, err := conn.ReadPartitions()
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	for _, p := range partition {
// 		if p.Topic == topic {
// 			return true
// 		}
// 	}

// 	return false
// }
