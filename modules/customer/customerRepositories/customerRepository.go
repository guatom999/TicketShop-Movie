package customerRepositories

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/guatom999/TicketShop-Movie/modules/customer"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	CustomerRepositoryService interface {
		InsertCustomer(pctx context.Context, req *customer.Customer) (primitive.ObjectID, error)
	}

	customerRepository struct {
		db *mongo.Client
	}
)

func NewCustomerRepository(db *mongo.Client) CustomerRepositoryService {

	return &customerRepository{db: db}

}

func (r *customerRepository) AddTicketCustomer(pctx context.Context) error {

	// ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	// defer cancel()

	// db := r.db.Database("movie_db")
	// col := db.Collection("movie")

	return nil

}

func (r *customerRepository) InsertCustomer(pctx context.Context, req *customer.Customer) (primitive.ObjectID, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("customer_db")
	col := db.Collection("customer")

	customerId, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: error insert one customer failed %s", err)
		return primitive.NilObjectID, errors.New("error: register failed")
	}

	return customerId.InsertedID.(primitive.ObjectID), nil
}
