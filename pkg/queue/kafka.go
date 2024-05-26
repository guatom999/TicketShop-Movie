package queue

import (
	"context"
	"fmt"

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

func KafkaReader(topic string) *kafka.Reader {
	// Define topic and consumer group
	// topic := "your-topic"
	groupID := "my-consumer-group"

	// Reader configuration with latest offset
	readerConfig := kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"}, // Adjust broker address(es)
		GroupID:     groupID,
		Topic:       topic,
		StartOffset: kafka.LastOffset,
	}

	// Create reader instance
	reader := kafka.NewReader(readerConfig)

	fmt.Println("Starting consumer...")

	// defer reader.Close()

	return reader

	// for {
	// 	// Read message batch with timeout
	// 	msg, err := reader.ReadMessage(context.Background())
	// 	if err != nil {
	// 		fmt.Println("Error reading message:", err)
	// 		break
	// 	}

	// 	// Process the message (e.g., extract key and value)
	// 	fmt.Printf("Received message: key=%s, value=%s\n", msg.Key, msg.Value)
	// }

	// fmt.Println("Consumer stopped.")

	// return nil
}

func IsTopicIsAlreadyExits(conn *kafka.Conn, topic string) bool {
	partition, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	for _, p := range partition {
		if p.Topic == topic {
			return true
		}
	}

	return false
}
