package inventoryRepositories

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/guatom999/TicketShop-Movie/modules/inventory"
	"github.com/guatom999/TicketShop-Movie/pkg/rest"
	"github.com/guatom999/TicketShop-Movie/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	InventoryRepositoryService interface {
		GetCustomerTicket(pctx context.Context, customerID string) ([]*inventory.CustomerTikcetRes, error)
		// FindLastCustomerTicket(pctx context.Context, customerId string, opts []*options.FindOneOptions) (*inventory.CustomerTikcetRes, error)
		// FindLastCustomerTicket(pctx context.Context, customerId string, opts any) (*inventory.CustomerTikcetRes, error)
		FindLastCustomerTicket(pctx context.Context, customerId string, opts []*options.FindOneOptions) (*inventory.CustomerTikcetRes, error)
		AddCustomerTicket(pctx context.Context, req *inventory.CustomerTicket) (primitive.ObjectID, error)
	}

	inventoryRepository struct {
		db *mongo.Client
	}
)

func NewInventoryRepository(db *mongo.Client) InventoryRepositoryService {
	return &inventoryRepository{db: db}
}

// func DbConnect(pctx context.Context, dbName string) *mongo.Collection {
// 	db := r.db.Database("inventory_db")
// }

// func (r *inventoryRepository) FindOneCustomerTicket(pctx context.Context, orderId)

func (r *inventoryRepository) FindOneTicketDetail(pctx context.Context, ticketId string) (*inventory.CustomerTikcetRes, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*10)
	defer cancel()

	db := r.db.Database("inventory_db")
	col := db.Collection("ticket_inventory")

	result := new(inventory.CustomerTicket)

	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertStringToObjectId(ticketId)}).Decode(result); err != nil {
		log.Printf("Error: find one ticket failed :%s", err.Error())
		return nil, err
	}

	return &inventory.CustomerTikcetRes{}, nil
}

func (r *inventoryRepository) findMovieDetail(movieId string) error {

	result, err := rest.ReqWithParams("http://localhost:8090/movie/getmovie/", movieId)
	if err != nil {
		log.Printf("Error: find one movie failed :%s", err.Error())
		return err
	}

	fmt.Println("Result:", result)

	return nil
}

func (r *inventoryRepository) GetCustomerTicket(pctx context.Context, customerID string) ([]*inventory.CustomerTikcetRes, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*30)
	defer cancel()

	db := r.db.Database("inventory_db")
	col := db.Collection("ticket_inventory")

	cur, err := col.Find(ctx, bson.M{"customer_id": customerID})
	if err != nil {
		log.Printf("Error: Failed to find Customer Ticket %s", err.Error())
		return make([]*inventory.CustomerTikcetRes, 0), nil
	}

	results := make([]*inventory.CustomerTikcetRes, 0)

	for cur.Next(ctx) {
		result := new(inventory.CustomerTicket)
		if err := cur.Decode(result); err != nil {
			log.Printf("Error: Failed to Decode From Finding %s", err.Error())
			return make([]*inventory.CustomerTikcetRes, 0), nil
		}

		results = append(results, &inventory.CustomerTikcetRes{
			TicketId:      result.Id.Hex(),
			OrderNumber:   result.OrderNumber,
			MovieId:       result.MovieId,
			MovieName:     result.MovieName,
			MovieImage:    result.PosterUrl,
			MovieDate:     result.MovieDate,
			MovieShowTime: result.MovieShowTime,
			Ticket_Image:  result.Ticket_Image,
			Created_At:    utils.GetStringTime(result.Created_At),
			Price:         result.Price,
			Seat:          result.Seat,
		})

	}

	return results, nil

}

// func (r *inventoryRepository) FindLastCustomerTicket(pctx context.Context, filter any, opts []*options.FindOneOptions) (*inventory.CustomerTikcetRes, error) {
func (r *inventoryRepository) FindLastCustomerTicket(pctx context.Context, customerId string, opts []*options.FindOneOptions) (*inventory.CustomerTikcetRes, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("inventory_db")
	col := db.Collection("ticket_inventory")

	// filter := bson.M{"customer_id": customerId}

	// sort := bson.M{"created_at": -1}

	// result := new(inventory.CustomerTikcetRes)
	result := new(inventory.CustomerTicket)

	if err := col.FindOne(ctx, bson.M{"customer_id": customerId}, opts...).Decode(result); err != nil {
		log.Printf("Error: Failed to Decode From Finding %s", err.Error())
		return nil, nil
	}

	return &inventory.CustomerTikcetRes{
		MovieId:      result.MovieId,
		MovieName:    result.MovieName,
		Ticket_Image: result.Ticket_Image,
		Created_At:   utils.GetStringTime(result.Created_At),
		Price:        result.Price / 100,
		Seat:         result.Seat,
	}, nil
}

func (r *inventoryRepository) AddCustomerTicket(pctx context.Context, req *inventory.CustomerTicket) (primitive.ObjectID, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("inventory_db")
	col := db.Collection("ticket_inventory")

	result, err := col.InsertOne(ctx, req)
	if err != nil {
		fmt.Printf("Error: Insert Customer Ticket Failed: %s", err.Error())
		return primitive.NilObjectID, errors.New("error: insert ticket failed")
	}

	return result.InsertedID.(primitive.ObjectID), nil
}
