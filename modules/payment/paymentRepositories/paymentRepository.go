package paymentRepositories

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/modules/payment"
	"github.com/guatom999/TicketShop-Movie/pkg/queue"
	"github.com/guatom999/TicketShop-Movie/utils"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/segmentio/kafka-go"
)

type (
	PaymentRepositoryService interface {
		// ReserveSeat(pctx context.Context) error
		ReserveSeat(pctx context.Context, cfg *config.Config, req *payment.MovieBuyReq) error
		// GetOffset(pctx context.Context) (int64 ,error)
		UpsertOfset(pctx context.Context, offset int64) error
	}

	paymentRepository struct {
		db *mongo.Client
	}
)

func NewPaymentRepository(db *mongo.Client) PaymentRepositoryService {
	return &paymentRepository{db: db}
}

func PaymentConsumer(pctx context.Context, cfg *config.Config, topic string) *kafka.Conn {
	conn := queue.KafkaConn(cfg)
	fmt.Println("kafka connect is success")

	topicConfigs := make([]kafka.TopicConfig, 0)

	if !queue.IsTopicIsAlreadyExits(conn, cfg.Kafka.Topic) {
		topicConfigs = append(topicConfigs, kafka.TopicConfig{
			Topic:             "buy",
			NumPartitions:     1,
			ReplicationFactor: 1,
		})
	}

	if err := conn.CreateTopics(topicConfigs...); err != nil {
		log.Printf("Erorr: Create Topic Failed %s", err.Error())
		panic(err.Error())
	}

	return conn

}

func (r *paymentRepository) UpsertOfset(pctx context.Context, offset int64) error {

	ctx, cancel := context.WithTimeout(pctx, time.Second*30)
	defer cancel()

	db := r.db.Database("payment_db")
	col := db.Collection("payment_queue")

	result, err := col.UpdateOne(ctx, bson.M{}, bson.M{"$set": bson.M{"offset": offset}}, options.Update().SetUpsert(true))
	if err != nil {
		log.Printf("Error: Update Upsert Offset Failed %s", err.Error())
		return errors.New("error: update offset failed")
	}

	fmt.Println("Result is", result)

	return nil

}

func (r *paymentRepository) ReserveSeat(pctx context.Context, cfg *config.Config, req *payment.MovieBuyReq) error {

	ctx, cancel := context.WithTimeout(pctx, time.Second*10)
	defer cancel()

	conn := PaymentConsumer(ctx, cfg, "buy")

	message := kafka.Message{
		Value: utils.EncodeMessage(req),
	}

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	_, err := conn.WriteMessages(message)
	if err != nil {
		log.Fatal("failed to write messages:", err)
		return errors.New("error: failed to send message")
	}

	// Close connection
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
		return errors.New("error: failed to close broker")
	}

	fmt.Println("Send Message Success")

	return nil
}
