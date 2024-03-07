package ticketRepositories

import (
	"context"
	"time"

	"github.com/guatom999/TicketShop-Movie/modules/ticket"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	TicketRepositoryService interface {
		AddCustomerTicket(pctx context.Context, req *ticket.Ticket) error
	}

	ticketRepository struct {
		db *mongo.Client
	}
)

func NewTicketRepository(db *mongo.Client) TicketRepositoryService {
	return &ticketRepository{
		db: db,
	}
}

func (r *ticketRepository) AddCustomerTicket(pctx context.Context, req *ticket.Ticket) error {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("tikcet_db")
	col := db.Collection("customer_ticket")

	col.InsertOne(ctx, req)

	return nil
}
