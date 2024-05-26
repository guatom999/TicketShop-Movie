package inventoryRepositories

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/guatom999/TicketShop-Movie/modules/inventory"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	InventoryRepositoryService interface {
	}

	inventoryRepository struct {
		db *mongo.Client
	}
)

func NewInventoryRepository(db *mongo.Client) InventoryRepositoryService {
	return &inventoryRepository{db: db}
}

func (r *inventoryRepository) FindCustomerTicket(pctx context.Context, customerID string) ([]*inventory.CustomerTicket, error) {

	_, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	return nil, nil

}

func (r *inventoryRepository) AddCustomerTicket(pctx context.Context, req *inventory.CustomerTicket) (primitive.ObjectID, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("inventory_db")
	col := db.Collection("ticket_inventory")

	result, err := col.InsertOne(ctx, req)
	if err != nil {
		fmt.Printf("Error: Insert Customer Ticket Failed: %s", err.Error())
		return primitive.NilObjectID, errors.New("error: insert ticket failed")
	}

	return result.InsertedID.(primitive.ObjectID), nil
}
