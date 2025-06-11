package ticketHandlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/guatom999/TicketShop-Movie/modules/ticket"
	"github.com/guatom999/TicketShop-Movie/modules/ticket/ticketUseCases"
	"github.com/labstack/echo/v4"
)

type (
	TicketHandlerService interface {
		AddCustomerTicket(c echo.Context) error
		FindCustomerTicket(c echo.Context) error
		HealthCheck(c echo.Context) error
	}

	ticketHandler struct {
		ticketUseCase ticketUseCases.TicketUseCaseService
	}
)

func NewTicketHandler(ticketUseCase ticketUseCases.TicketUseCaseService) TicketHandlerService {
	return &ticketHandler{ticketUseCase: ticketUseCase}
}

func (h *ticketHandler) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "status ok")
}

func (h *ticketHandler) AddCustomerTicket(c echo.Context) error {

	ctx := context.Background()

	req := new(ticket.AddTikcetReq)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "bad req")
	}

	result, err := h.ticketUseCase.AddCustomerTicket(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("add ticket success %s", result.Hex()))
}

func (h *ticketHandler) FindCustomerTicket(c echo.Context) error {

	ctx := context.Background()

	customerId := c.Param("customer_id")

	result, err := h.ticketUseCase.FindCustomerTicket(ctx, customerId)
	if err != nil {
		return c.JSON(http.StatusBadGateway, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
