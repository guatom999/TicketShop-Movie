package bookRepositories

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	BookRepositoryService interface {
	}

	bookRepository struct {
		db *mongo.Client
	}
)

func NewBookRepository() BookRepositoryService {
	return &bookRepository{}
}

func (r *bookRepository) ConnectBookingDb() *mongo.Database {
	return r.db.Database("booking_db")
}

func (r *bookRepository) BuyTicket(pctx context.Context) {

}
