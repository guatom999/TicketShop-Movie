package inventoryRepositories

import (
	"context"
	"time"

	"github.com/guatom999/TicketShop-Movie/modules/inventory"
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

func (r *inventoryRepository) FindCustomerTicket(pctx context.Context, customerID string) ([]*inventory.Inventory, error) {

	_, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	return nil, nil

}
