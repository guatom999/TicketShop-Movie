package ticketUseCases

import (
	"context"

	"github.com/guatom999/TicketShop-Movie/modules/ticket"
	"github.com/guatom999/TicketShop-Movie/modules/ticket/ticketRepositories"
	"github.com/guatom999/TicketShop-Movie/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	TicketUseCaseService interface {
		AddCustomerTicket(pctx context.Context, req *ticket.AddTikcetReq) (primitive.ObjectID, error)
		FindCustomerTicket(pctx context.Context, customerId string) ([]*ticket.TicketShowCase, error)
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
		MovieName:  req.MovieName,
		Seat:       req.Seat,
		Date:       req.Date,
		Time:       req.Time,
		Price:      req.Price,
		CreatedAt:  utils.GetLocaltime(),
		UpdatedAt:  utils.GetLocaltime(),
	})
	if err != nil {
		return primitive.NilObjectID, nil
	}

	return result, nil

}

func (u *ticketUseCase) FindCustomerTicket(pctx context.Context, customerId string) ([]*ticket.TicketShowCase, error) {

	tickets, err := u.ticketRepo.FindTicket(pctx, customerId)

	if err != nil {
		return tickets, err
	}

	return tickets, nil

}
