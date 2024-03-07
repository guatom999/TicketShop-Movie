package ticketRepositories

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/guatom999/TicketShop-Movie/modules/ticket"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	TicketRepositoryService interface {
		AddCustomerTicket(pctx context.Context, req *ticket.Ticket) (primitive.ObjectID, error)
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

func (r *ticketRepository) AddCustomerTicket(pctx context.Context, req *ticket.Ticket) (primitive.ObjectID, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("tikcet_db")
	col := db.Collection("customer_ticket")

	result, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: Add Ticket Failed: %v", err)
		return primitive.NilObjectID, errors.New("error: add tikcet failed")
	}

	log.Println("insert object id is", result.InsertedID.(primitive.ObjectID))

	return result.InsertedID.(primitive.ObjectID), nil
}
