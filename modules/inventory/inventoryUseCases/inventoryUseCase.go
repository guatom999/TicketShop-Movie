package inventoryUseCases

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/guatom999/TicketShop-Movie/modules/inventory"
	"github.com/guatom999/TicketShop-Movie/modules/inventory/inventoryRepositories"
	"github.com/guatom999/TicketShop-Movie/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	InventoryUseCaseService interface {
		GetCustomerTicket(pctx context.Context, customerId string) ([]*inventory.CustomerTikcetRes, error)
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

func (u *inventoryUseCase) GetCustomerTicket(pctx context.Context, customerId string) ([]*inventory.CustomerTikcetRes, error) {

	results, err := u.inventoryRepo.GetCustomerTicket(pctx, strings.TrimPrefix(customerId, "customer:"))
	if err != nil {
		log.Printf("Error: get movie show time failed %s", err.Error())
		return nil, err
	}

	// _ = u.inventoryRepo.GetMovieDetails(pctx, results[0].MovieId)

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

	fmt.Println("MovieDate in inventory time is :::::::::::::>", req.MovieDate, "MovieShowTime in inventory time is", req.MovieShowTime)

	insertId, err := u.inventoryRepo.AddCustomerTicket(pctx, &inventory.CustomerTicket{
		CustomerId:    req.CustomerId,
		Ticket_Image:  req.TicketUrl,
		MovieId:       req.MovieId,
		MovieName:     req.MovieName,
		MovieDate:     req.MovieDate,
		MovieShowTime: req.MovieShowTime,
		PosterUrl:     req.PosterImage,
		OrderNumber:   req.OrderNumber,
		Created_At:    utils.GetLocaltime(),
		Price:         req.Price,
		Seat:          req.SeatNo,
	})

	if err != nil {
		fmt.Println("insertId ", insertId)
	}

	fmt.Println("insertID is", insertId)

}
