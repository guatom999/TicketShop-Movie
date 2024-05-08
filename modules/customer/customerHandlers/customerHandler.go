package customerHandlers

import (
	"context"
	"net/http"

	"github.com/guatom999/TicketShop-Movie/modules/customer"
	"github.com/guatom999/TicketShop-Movie/modules/customer/customerUseCases"
	"github.com/labstack/echo/v4"
)

type (
	CustomerHandlerService interface {
		Login(c echo.Context) error
		Register(c echo.Context) error
	}

	customerHandler struct {
		customerUseCase customerUseCases.CustomerUseCaseService
	}
)

func NewCustomerHandler(customerUseCase customerUseCases.CustomerUseCaseService) CustomerHandlerService {
	return &customerHandler{customerUseCase: customerUseCase}
}

func (h *customerHandler) Login(c echo.Context) error {

	ctx := context.Background()

	req := new(customer.LoginReq)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "something weng wrong")
	}

	res, err := h.customerUseCase.Login(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (h *customerHandler) Register(c echo.Context) error {

	ctx := context.Background()

	req := new(customer.RegisterReq)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := h.customerUseCase.Register(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}
