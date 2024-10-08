package inventoryUseCases

import (
	"context"
	"fmt"
	"log"

	"github.com/guatom999/TicketShop-Movie/modules/inventory"
	"github.com/guatom999/TicketShop-Movie/modules/inventory/inventoryRepositories"
	"github.com/guatom999/TicketShop-Movie/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	InventoryUseCaseService interface {
		FindCustomerTicket(pctx context.Context, customerId string) ([]*inventory.CustomerTikcetRes, error)
		FindLastCustomerTicket(pctx context.Context, customerId string) (*inventory.CustomerTikcetRes, error)
		AddCustomerTicket(pctx context.Context, req *inventory.AddCustomerTicketReq)
	}

	inventoryUseCase struct {
		inventoryRepo inventoryRepositories.InventoryRepositoryService
	}
)

func NewInventoryUseCase(inventoryRepo inventoryRepositories.InventoryRepositoryService) InventoryUseCaseService {
	return &inventoryUseCase{inventoryRepo: inventoryRepo}
}

func (u *inventoryUseCase) FindCustomerTicket(pctx context.Context, customerId string) ([]*inventory.CustomerTikcetRes, error) {

	results, err := u.inventoryRepo.FindCustomerTicket(pctx, customerId)
	if err != nil {
		log.Printf("Error: get movie show time failed %s", err.Error())
		return nil, err
	}

	return results, nil

}

func (u *inventoryUseCase) FindLastCustomerTicket(pctx context.Context, customerId string) (*inventory.CustomerTikcetRes, error) {

	// filter := bson.D{}
	// filterOptions := make([]*options.FindOneOptions, 0)

	findItemOption := make([]*options.FindOneOptions, 0)

	findItemOption = append(findItemOption, options.FindOne().SetSort(bson.D{{"created_at", -1}}))

	// filter = append(filter, options.FindOneOptions().WithFilter(filter))
	// filterOptions = append(filterOptions,bson.D{{"created_at":-1}})

	result, err := u.inventoryRepo.FindLastCustomerTicket(pctx, customerId, findItemOption)
	if err != nil {
		log.Printf("Error: get movie show time failed %s", err.Error())
		return nil, err
	}

	return result, nil
}

func (u *inventoryUseCase) AddCustomerTicket(pctx context.Context, req *inventory.AddCustomerTicketReq) {

	insertId, err := u.inventoryRepo.AddCustomerTicket(pctx, &inventory.CustomerTicket{
		CustomerId:   req.CustomerId,
		MovieId:      req.MovieId,
		Ticket_Image: req.TicketUrl,
		MovieName:    req.MovieName,
		Created_At:   utils.GetLocaltime(),
		Price:        req.Quantity * 150,
		Seat:         req.SeatNo,
	})

	if err != nil {
		fmt.Println("insertId ", insertId)
	}

	fmt.Println("insertID is", insertId)

}
