package server

import (
	"github.com/guatom999/TicketShop-Movie/modules/ticket/ticketHandlers"
	"github.com/guatom999/TicketShop-Movie/modules/ticket/ticketRepositories"
	"github.com/guatom999/TicketShop-Movie/modules/ticket/ticketUseCases"
)

func (s *server) TicketModule() {
	ticketRepo := ticketRepositories.NewTicketRepository(s.db)
	ticketUseCase := ticketUseCases.NewTicketUseCase(ticketRepo)
	ticketHandler := ticketHandlers.NewTicketHandler(ticketUseCase)

	tikcetRouter := s.app.Group("/ticket")

	tikcetRouter.POST("/add", ticketHandler.AddCustomerTicket)

}
