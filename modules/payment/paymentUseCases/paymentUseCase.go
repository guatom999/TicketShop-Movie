package paymentUseCases

import (
	"context"

	"github.com/guatom999/TicketShop-Movie/modules/payment/paymentRepositories"
)

type (
	PaymentUseCaseService interface {
	}

	paymentUseCase struct {
		paymentRepo paymentRepositories.PaymentRepositoryService
	}
)

func NewPaymentUseCase(paymentRepo paymentRepositories.PaymentRepositoryService) PaymentUseCaseService {
	return &paymentUseCase{paymentRepo: paymentRepo}
}

func (u *paymentUseCase) BuyTicket(pctx context.Context) error {

	return nil

}
