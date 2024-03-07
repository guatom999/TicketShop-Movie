package ticketHandlers

import (
	"net/http"

	"github.com/guatom999/TicketShop-Movie/modules/ticket"
	"github.com/guatom999/TicketShop-Movie/modules/ticket/ticketUseCases"
	"github.com/labstack/echo/v4"
)

type (
	TicketHandlerService interface {
	}

	ticketHandler struct {
		ticketUseCase ticketUseCases.TicketUseCaseService
	}
)

func NewTicketHandler(ticketUseCase ticketUseCases.TicketUseCaseService) TicketHandlerService {
	return &ticketHandler{ticketUseCase: ticketUseCase}
}

func (h *ticketHandler) AddCustomerTicket(c echo.Context) error {
	req := new(ticket.AddTikcetReq)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "bad req")
	}

	return c.JSON(http.StatusOK, "Add Customer Ticket Successful")
}
