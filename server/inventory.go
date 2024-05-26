package server

import (
	"github.com/guatom999/TicketShop-Movie/modules/inventory/inventoryHandlers"
	"github.com/guatom999/TicketShop-Movie/modules/inventory/inventoryRepositories"
	"github.com/guatom999/TicketShop-Movie/modules/inventory/inventoryUseCases"
)

func (s *server) InventoryModule() {
	inventoryRepo := inventoryRepositories.NewInventoryRepository(s.db)
	inventoryUseCase := inventoryUseCases.NewInventoryUseCase(inventoryRepo)
	inventoryHandler := inventoryHandlers.NewInventoryHandler(inventoryUseCase)
	inventoryQueueHandler := inventoryHandlers.NewInventoryQueueHandler(s.cfg, inventoryUseCase)

	ticketRouter := s.app.Group("/inventory")

	go inventoryQueueHandler.AddCustomerTransaction()

	// tikcetRouter.POST("/add", inventoryHandler.FindCustomerTicket)
	ticketRouter.GET("/:customerid", inventoryHandler.FindCustomerTicket)
	// ticketRouter.PO

}
