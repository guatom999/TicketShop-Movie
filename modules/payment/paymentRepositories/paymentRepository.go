package paymentRepositories

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/modules/model"
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
		ReserveSeat(pctx context.Context, cfg *config.Config, req *payment.ReserveSeatReq) error
		AddTicketToCustomer(pctx context.Context, cfg *config.Config, req *payment.AddCustomerTicket) error
		GetOffset(pctx context.Context) (int64, error)
		UpsertOfset(pctx context.Context, offset int64) error
		RollBackReserveSeat(pctx context.Context, cfg *config.Config, req *payment.RollBackReservedSeatReq) error
	}

	paymentRepository struct {
		// cfg *config.Config
		db *mongo.Client
	}
)

func NewPaymentRepository(db *mongo.Client) PaymentRepositoryService {
	return &paymentRepository{db: db}
}

func PaymentProducer(pctx context.Context, cfg *config.Config, topic string) *kafka.Conn {
	conn := queue.KafkaConn(cfg, topic)

	// topicConfigs := []kafka.TopicConfig{
	// 	{
	// 		Topic:             "buy",
	// 		NumPartitions:     1,
	// 		ReplicationFactor: 1,
	// 	},
	// 	{
	// 		Topic:             "rollback",
	// 		NumPartitions:     1,
	// 		ReplicationFactor: 1,
	// 	},
	// }

	topicConfigs := kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     1,
		ReplicationFactor: 1,
	}

	// topicConfigs := make([]kafka.TopicConfig, 0)

	// if !queue.IsTopicIsAlreadyExits(conn, cfg.Kafka.Topic) {
	// 	topicConfigs = append(topicConfigs, kafka.TopicConfig{
	// 		Topic:             topic,
	// 		NumPartitions:     1,
	// 		ReplicationFactor: 1,
	// 	})
	// }

	if err := conn.CreateTopics(topicConfigs); err != nil {
		log.Printf("Erorr: Create Topic Failed %s", err.Error())
		panic(err.Error())
	}

	return conn

}

func (r *paymentRepository) GetOffset(pctx context.Context) (int64, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*30)
	defer cancel()

	db := r.db.Database("payment_db")
	col := db.Collection("payment_queue")

	result := new(model.KafKaOffset)

	if err := col.FindOne(ctx, bson.M{}).Decode(result); err != nil {
		log.Printf("Error: GetOffset  failed: %s", err.Error())
		return -1, errors.New("error: getoffset failed")
	}

	return 0, nil
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

func (r *paymentRepository) ReserveSeat(pctx context.Context, cfg *config.Config, req *payment.ReserveSeatReq) error {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	// conn := PaymentProducer(ctx, cfg, "buy-ticket")
	conn := PaymentProducer(ctx, cfg, "reserve-seat")

	message := kafka.Message{
		Key:   []byte("movie"),
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

func (r *paymentRepository) RollBackReserveSeat(pctx context.Context, cfg *config.Config, req *payment.RollBackReservedSeatReq) error {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	conn := PaymentProducer(ctx, cfg, "rollback")

	message := kafka.Message{
		Key:   []byte("movie"),
		Value: utils.EncodeMessage(req),
	}

	conn.SetReadDeadline(time.Now().Add(time.Second * 20))
	_, err := conn.WriteMessages(message)

	if err != nil {
		log.Fatalf("Error writing message: %v", err)
		return errors.New("error: write message failed")
	}

	if err := conn.Close(); err != nil {
		log.Fatalf("Error Failed to close queue: %v", err)
		return errors.New("error: failed to close message queue")
	}
	fmt.Println("Send Message Success")

	return nil
}

func (r *paymentRepository) AddTicketToCustomer(pctx context.Context, cfg *config.Config, req *payment.AddCustomerTicket) error {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	conn := PaymentProducer(ctx, cfg, "add-ticket")

	message := kafka.Message{
		Value: utils.EncodeMessage(req),
	}

	conn.SetReadDeadline(time.Now().Add(time.Second * 20))
	_, err := conn.WriteMessages(message)

	if err != nil {
		log.Fatalf("Error writing message: %v", err)
		return errors.New("error: write message failed")
	}

	if err := conn.Close(); err != nil {
		log.Fatalf("Error Failed to close queue: %v", err)
		return errors.New("error: failed to close messagq queue")
	}
	fmt.Println("Send Message Success")

	return nil
}
