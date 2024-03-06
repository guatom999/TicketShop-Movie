package ticketRepositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	TicketRepositoryService interface {
	}

	ticketRepository struct {
		db *mongo.Client
	}
)

func NewTicketRepository() TicketRepositoryService {
	return ticketRepository{}
}

func (r *ticketRepository) AddCustomerTicket(pctx context.Context, req any) error {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("tikcet_db")
	col := db.Collection("customer_ticket")

	col.InsertOne(ctx, req)

	return nil
}
