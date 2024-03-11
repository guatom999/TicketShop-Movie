package inventoryRepositories

import (
	"context"
	"time"

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

	_, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	return nil, nil

}
