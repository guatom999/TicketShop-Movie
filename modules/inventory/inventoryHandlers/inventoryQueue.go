package inventoryHandlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/modules/inventory"
	"github.com/guatom999/TicketShop-Movie/modules/inventory/inventoryUseCases"
	"github.com/guatom999/TicketShop-Movie/pkg/queue"
)

type (
	InventoryQueueHanlderService interface {
		AddCustomerTransaction()
	}

	inventoryQueueHanlder struct {
		cfg              *config.Config
		inventoryUseCase inventoryUseCases.InventoryUseCaseService
	}
)

func NewInventoryQueueHandler(cfg *config.Config, inventoryUseCase inventoryUseCases.InventoryUseCaseService) InventoryQueueHanlderService {
	return &inventoryQueueHanlder{
		cfg:              cfg,
		inventoryUseCase: inventoryUseCase,
	}
}

func (h *inventoryQueueHanlder) AddCustomerTransaction() {

	ctx := context.Background()

	data := new(inventory.AddCustomerTicketReq)

	reader := queue.KafkaReader("add-ticket", "inventory-group")
	defer reader.Close()

	for {

		message, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("Error reading message: %s", err.Error())
			break
		}

		if err := json.Unmarshal(message.Value, data); err != nil {
			fmt.Printf("Error: Unmarshal error %s", err.Error())
		}

		h.inventoryUseCase.AddCustomerTicket(ctx, data)

	}

}
