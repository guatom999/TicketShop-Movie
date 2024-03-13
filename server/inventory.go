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

	tikcetRouter := s.app.Group("/inventory")

	// tikcetRouter.POST("/add", inventoryHandler.FindCustomerTicket)
	tikcetRouter.GET("/:customerid", inventoryHandler.FindCustomerTicket)

}
