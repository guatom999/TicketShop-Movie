package inventoryUseCases

import (
	"context"
	"fmt"
	"log"

	"github.com/guatom999/TicketShop-Movie/modules/inventory"
	"github.com/guatom999/TicketShop-Movie/modules/inventory/inventoryRepositories"
	"github.com/guatom999/TicketShop-Movie/pkg/rest"
)

type (
	InventoryUseCaseService interface {
		FindCustomerTicket(pctx context.Context, customerId string) ([]*inventory.CustomerInventoryRes, error)
	}

	inventoryUseCase struct {
		inventoryRepo inventoryRepositories.InventoryRepositoryService
	}
)

func NewInventoryUseCase(inventoryRepo inventoryRepositories.InventoryRepositoryService) InventoryUseCaseService {
	return &inventoryUseCase{inventoryRepo: inventoryRepo}
}

func (u *inventoryUseCase) FindCustomerTicket(pctx context.Context, customerId string) ([]*inventory.CustomerInventoryRes, error) {

	fmt.Println("Find CustomerTicket")
	baseUrl := "http://localhost:8090/movie/getmovieshowtime/"
	paramsUrl := baseUrl + customerId

	result, err := rest.Request(paramsUrl)
	if err != nil {
		log.Printf("Error: get movie show time failed %s", err.Error())
		return nil, err
	}

	log.Println("Result is", result)

	return nil, nil

}
