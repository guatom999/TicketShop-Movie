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
		// CheckOutWithCreditCard(c echo.Context) error
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

	req := new(payment.MovieBuyReq)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.paymentUseCase.BuyTicket(ctx, h.cfg, req); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "Buy Ticket Success")
}
