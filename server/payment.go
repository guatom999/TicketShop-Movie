package server

import (
	paymentHandler "github.com/guatom999/TicketShop-Movie/modules/payment/paymentHandlers"
	"github.com/guatom999/TicketShop-Movie/modules/payment/paymentRepositories"
	"github.com/guatom999/TicketShop-Movie/modules/payment/paymentUseCases"
)

func (s *server) PaymentModules() {
	paymentRepo := paymentRepositories.NewPaymentRepository(s.db)
	paymentUseCase := paymentUseCases.NewPaymentUseCase(paymentRepo)
	paymentHandler := paymentHandler.NewPaymentHanlder(s.cfg, paymentUseCase)

	// _ = paymentHandler

	router := s.app.Group("/payment")

	router.POST("/buyticket", paymentHandler.BuyTicket)

}
