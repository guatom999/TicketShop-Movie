package paymentHandler

import (
	"context"
	"net/http"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/modules/payment"
	"github.com/guatom999/TicketShop-Movie/modules/payment/paymentUseCases"
	"github.com/labstack/echo/v4"
)

type (
	PaymentHandlerService interface {
		BuyTicket(c echo.Context) error
	}

	paymentHandler struct {
		paymentUseCase paymentUseCases.PaymentUseCaseService
		cfg            *config.Config
	}
)

func NewPaymentHanlder(cfg *config.Config, paymentUseCase paymentUseCases.PaymentUseCaseService) PaymentHandlerService {
	return &paymentHandler{
		cfg:            cfg,
		paymentUseCase: paymentUseCase,
	}
}

func (h *paymentHandler) BuyTicket(c echo.Context) error {

	ctx := context.Background()

	_ = ctx

	req := new(payment.MovieBuyReq)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return nil
}
