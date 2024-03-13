package ticketRepositories

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/guatom999/TicketShop-Movie/modules/ticket"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	TicketRepositoryService interface {
		AddTicket(pctx context.Context, req *ticket.Ticket) (primitive.ObjectID, error)
		FindTicket(pctx context.Context, customerId string) ([]*ticket.TicketShowCase, error)
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

func (r *ticketRepository) AddTicket(pctx context.Context, req *ticket.Ticket) (primitive.ObjectID, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("ticket_db")
	col := db.Collection("customer_ticket")

	result, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: Add Ticket Failed: %v", err)
		return primitive.NilObjectID, errors.New("error: add tikcet failed")
	}

	log.Println("insert object id is", result.InsertedID.(primitive.ObjectID))

	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *ticketRepository) FindTicket(pctx context.Context, customerId string) ([]*ticket.TicketShowCase, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("ticket_db")
	col := db.Collection("customer_ticket")

	cur, err := col.Find(ctx, bson.M{"customer_id": customerId})
	if err != nil {
		log.Printf("Error: FindTicket Failed %s", err.Error())
		return make([]*ticket.TicketShowCase, 0), errors.New("error: find ticket failed")
	}

	tickets := make([]*ticket.TicketShowCase, 0)

	for cur.Next(ctx) {
		result := new(ticket.Ticket)
		if err := cur.Decode(result); err != nil {
			log.Printf("Error: Failed to Decoed Ticket %s", err.Error())
			return make([]*ticket.TicketShowCase, 0), err
		}

		tickets = append(tickets, &ticket.TicketShowCase{
			Title: result.MovieName,
			Seat:  result.Seat,
			Date:  result.Date,
			Time:  result.Time,
			Price: result.Price,
		})
	}

	return tickets, nil
}
