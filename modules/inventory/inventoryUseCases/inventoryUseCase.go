package inventoryUseCases

import (
	"context"
	"fmt"
	"log"

	"github.com/guatom999/TicketShop-Movie/modules/inventory"
	"github.com/guatom999/TicketShop-Movie/modules/inventory/inventoryRepositories"
	"github.com/guatom999/TicketShop-Movie/pkg/rest"
	"github.com/guatom999/TicketShop-Movie/utils"
)

type (
	InventoryUseCaseService interface {
		FindCustomerTicket(pctx context.Context, customerId string) ([]*inventory.AddCustomerTicketReq, error)
		AddCustomerTicket(pctx context.Context, req *inventory.AddCustomerTicketReq)
	}

	inventoryUseCase struct {
		inventoryRepo inventoryRepositories.InventoryRepositoryService
	}
)

func NewInventoryUseCase(inventoryRepo inventoryRepositories.InventoryRepositoryService) InventoryUseCaseService {
	return &inventoryUseCase{inventoryRepo: inventoryRepo}
}

func (u *inventoryUseCase) FindCustomerTicket(pctx context.Context, customerId string) ([]*inventory.AddCustomerTicketReq, error) {

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

func (u *inventoryUseCase) AddCustomerTicket(pctx context.Context, req *inventory.AddCustomerTicketReq) {

	insertId, err := u.inventoryRepo.AddCustomerTicket(pctx, &inventory.CustomerTicket{
		CustomerId: req.CustomerId,
		MovieId:    req.MovieId,
		MovieName:  "Test",
		Created_At: utils.GetLocaltime(),
		Price:      req.Quantity * 150,
		Seat:       req.SeatNo,
	})

	if err != nil {
		fmt.Println("insertId ", insertId)
	}

	fmt.Println("insertID is", insertId)

}
