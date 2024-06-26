package server

import (
	paymentHandler "github.com/guatom999/TicketShop-Movie/modules/payment/paymentHandlers"
	"github.com/guatom999/TicketShop-Movie/modules/payment/paymentRepositories"
	"github.com/guatom999/TicketShop-Movie/modules/payment/paymentUseCases"
)

func (s *server) PaymentModule() {
	paymentRepo := paymentRepositories.NewPaymentRepository(s.db)
	paymentUseCase := paymentUseCases.NewPaymentUseCase(paymentRepo, s.cfg, s.omise)
	paymentHandler := paymentHandler.NewPaymentHanlder(s.cfg, paymentUseCase)

	// _ = paymentHandler

	router := s.app.Group("/payment")

	router.POST("/buyticket", paymentHandler.BuyTicket)
	// router.POST("/checkoutwithcreditcard", paymentHandler.CheckOutWithCreditCard)

}
