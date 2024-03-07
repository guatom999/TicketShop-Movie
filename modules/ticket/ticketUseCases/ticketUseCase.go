package ticketUseCases

import (
	"context"

	"github.com/guatom999/TicketShop-Movie/modules/ticket"
	"github.com/guatom999/TicketShop-Movie/modules/ticket/ticketRepositories"
)

type (
	TicketUseCaseService interface {
		AddCustomerTicket(pctx context.Context, req *ticket.AddTikcetReq) error
	}

	ticketUseCase struct {
		ticketRepo ticketRepositories.TicketRepositoryService
	}
)

func NewTicketUseCase(ticketRepo ticketRepositories.TicketRepositoryService) TicketUseCaseService {
	return &ticketUseCase{
		ticketRepo: ticketRepo,
	}
}

func (u *ticketUseCase) AddCustomerTicket(pctx context.Context, req *ticket.AddTikcetReq) error {

	result, err := u.ticketRepo.AddCustomerTicket(pctx, &ticket.Ticket{
		MovieId:    req.MovidId,
		CustomerId: req.CustomerId,
	})
	if err != nil {
		return nil
	}

	_ = result

	return nil

}
