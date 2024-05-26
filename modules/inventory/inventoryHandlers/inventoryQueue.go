package inventoryHandlers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/modules/inventory/inventoryUseCases"
	"github.com/guatom999/TicketShop-Movie/modules/movie"
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

	data := new(movie.ReserveSeatReqTest)

	reader := queue.KafkaReader("buy-ticket")
	defer reader.Close()

	for {

		message, err := reader.ReadMessage(ctx)
		fmt.Println("inventory message is =================================>", message)
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		fmt.Println("message is", message)

		if err := json.Unmarshal(message.Value, data); err != nil {
			fmt.Printf("Error: Unmarshal error %s", err.Error())
		}

	}

}
