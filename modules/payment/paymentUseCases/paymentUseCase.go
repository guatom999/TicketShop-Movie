package paymentUseCases

import (
	"context"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/modules/payment"
	"github.com/guatom999/TicketShop-Movie/modules/payment/paymentRepositories"
)

type (
	PaymentUseCaseService interface {
		BuyTicket(pctx context.Context, cfg *config.Config, req *payment.MovieBuyReq) error
	}

	paymentUseCase struct {
		paymentRepo paymentRepositories.PaymentRepositoryService
	}
)

func NewPaymentUseCase(paymentRepo paymentRepositories.PaymentRepositoryService) PaymentUseCaseService {
	return &paymentUseCase{paymentRepo: paymentRepo}
}

func (u *paymentUseCase) BuyTicket(pctx context.Context, cfg *config.Config, req *payment.MovieBuyReq) error {

	if err := u.paymentRepo.ReserveSeat(pctx, cfg, req); err != nil {
		return err
	}

	return nil

}
