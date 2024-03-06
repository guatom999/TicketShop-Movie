package inventoryRepositories

import (
	"context"

	"github.com/guatom999/TicketShop-Movie/modules/inventory"
)

type (
	InventoryRepositoryService interface {
	}

	inventoryRepository struct {
	}
)

func NewInventoryRepository() InventoryRepositoryService {
	return &inventoryRepository{}
}

func (r *inventoryRepository) FindCustomerTicket(pctx context.Context, customerID string) ([]*inventory.Inventory, error) {

	// ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	// defer cancel()

	// db := r.db.Database("inventory_db")
	// col := db.Collection("ticket_inventory")

	return nil, nil

}
