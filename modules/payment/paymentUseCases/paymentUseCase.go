package paymentUseCases

import (
	"context"
	"errors"
	"log"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/modules/payment"
	"github.com/guatom999/TicketShop-Movie/modules/payment/paymentRepositories"
	"github.com/guatom999/TicketShop-Movie/pkg/queue"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

type (
	PaymentUseCaseService interface {
		BuyTicket(pctx context.Context, cfg *config.Config, req *payment.MovieBuyReq) error
		CheckOutWithCreditCard(req *payment.CheckOutWithCreditCard) error
	}

	paymentUseCase struct {
		paymentRepo paymentRepositories.PaymentRepositoryService
		cfg         *config.Config
		opnClient   *omise.Client
	}
)

func NewPaymentUseCase(paymentRepo paymentRepositories.PaymentRepositoryService, cfg *config.Config, opnClient *omise.Client) PaymentUseCaseService {
	return &paymentUseCase{
		paymentRepo: paymentRepo,
		cfg:         cfg,
		opnClient:   opnClient,
	}
}

func (u *paymentUseCase) CheckOutWithCreditCard(req *payment.CheckOutWithCreditCard) error {

	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:   req.Price,
		Currency: "thb",
		Card:     req.Token,
	}
	if err := u.opnClient.Do(charge, createCharge); err != nil {
		log.Fatalf("omise clinet failed  %s", err.Error())
		return errors.New("error: omise client failed")
	}

	log.Printf("charge: %s  amount: %s %d\n", charge.ID, charge.Currency, charge.Amount)

	return nil
}

// func PaymentConsumer(pctx context.Context, cfg *config.Config, topic string) *kafka.Conn {
// 	conn := queue.KafkaConn(cfg)
// 	fmt.Println("kafka connect is success")

// 	topicConfigs := make([]kafka.TopicConfig, 0)

// 	if !queue.IsTopicIsAlreadyExits(conn, cfg.Kafka.Topic) {
// 		topicConfigs = append(topicConfigs, kafka.TopicConfig{
// 			Topic:             "buy",
// 			NumPartitions:     1,
// 			ReplicationFactor: 1,
// 		})
// 	}

// 	if err := conn.CreateTopics(topicConfigs...); err != nil {
// 		log.Printf("Erorr: Create Topic Failed %s", err.Error())
// 		panic(err.Error())
// 	}

// 	return conn

// }

func BuyTicketConsumer(pctx context.Context, topic string, resCh chan *payment.PaymentReserveRes) {

	reader := queue.KafkaReader("buy-ticket")
	defer reader.Close()

}

func (u *paymentUseCase) BuyTicket(pctx context.Context, cfg *config.Config, req *payment.MovieBuyReq) error {

	if err := u.CheckOutWithCreditCard(&payment.CheckOutWithCreditCard{Token: req.Token, Price: req.Price}); err != nil {
		return err
	}

	if err := u.paymentRepo.ReserveSeat(pctx, cfg, &payment.ReserveSeatReq{
		MovieId: req.MovieId,
		SeatNo:  req.SeatNo,
	}); err != nil {
		return err
	}

	return nil

}
