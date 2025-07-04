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
		TestUpload(c echo.Context) error
		HealthCheck(c echo.Context) error
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

func (h *paymentHandler) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "status ok")
}

func (h *paymentHandler) BuyTicket(c echo.Context) error {

	ctx := context.Background()

	req := new(payment.MovieBuyReq)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := h.paymentUseCase.BuyTicket(ctx, h.cfg, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)
}

func (h *paymentHandler) TestUpload(c echo.Context) error {

	// file, err := c.FormFile("image")
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, err.Error())
	// }

	// blobfile, err := file.Open()
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, err.Error())
	// }

	// if err := h.paymentUseCase.UploadFileTest(blobfile, ""); err != nil {
	// 	return c.JSON(http.StatusInternalServerError, err)
	// }

	return c.JSON(http.StatusOK, "test")
}
