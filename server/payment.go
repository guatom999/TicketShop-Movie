package server

import (
	paymentHandler "github.com/guatom999/TicketShop-Movie/modules/payment/paymentHandlers"
	"github.com/guatom999/TicketShop-Movie/modules/payment/paymentRepositories"
	"github.com/guatom999/TicketShop-Movie/modules/payment/paymentUseCases"
)

func (s *server) PaymentModule() {

	// client, err := storage.NewClient(context.Background())
	// if err != nil {
	// 	log.Fatalf("Failed to create client: %v", err)
	// }

	paymentRepo := paymentRepositories.NewPaymentRepository(s.db)
	// paymentUseCase := paymentUseCases.NewPaymentUseCase(paymentRepo, s.cfg, s.omise, client)
	paymentUseCase := paymentUseCases.NewPaymentUseCase(paymentRepo, s.cfg, s.omise)
	paymentHandler := paymentHandler.NewPaymentHanlder(s.cfg, paymentUseCase)

	// ctx := context.Background()

	// resCh := make(chan *payment.RollBackReserveSeatRes)

	// go paymentUseCase.BuyTicketConsumer(ctx, "rollback", resCh)

	// _ = paymentHandler

	router := s.app.Group("/payment")

	router.POST("/buyticket", paymentHandler.BuyTicket)
	router.POST("/testupload", paymentHandler.TestUpload)
	// router.POST("/checkoutwithcreditcard", paymentHandler.CheckOutWithCreditCard)

}
