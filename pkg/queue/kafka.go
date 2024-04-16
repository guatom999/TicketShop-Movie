package queue

import (
	"context"

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
