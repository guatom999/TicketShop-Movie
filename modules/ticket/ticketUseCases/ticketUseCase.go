package ticketUseCases

import (
	"context"

	"github.com/guatom999/TicketShop-Movie/modules/ticket"
	"github.com/guatom999/TicketShop-Movie/modules/ticket/ticketRepositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	TicketUseCaseService interface {
		AddCustomerTicket(pctx context.Context, req *ticket.AddTikcetReq) (primitive.ObjectID, error)
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

func (u *ticketUseCase) AddCustomerTicket(pctx context.Context, req *ticket.AddTikcetReq) (primitive.ObjectID, error) {

	result, err := u.ticketRepo.AddTicket(pctx, &ticket.Ticket{
		MovieId:    req.MovieId,
		CustomerId: req.CustomerId,
		Seat:       req.Seat,
	})
	if err != nil {
		return primitive.NilObjectID, nil
	}

	return result, nil

}
