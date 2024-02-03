package migrate

import (
	"context"
	"log"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/database"
	"github.com/guatom999/TicketShop-Movie/modules/book"
	"github.com/guatom999/TicketShop-Movie/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func BookingMigrate(pctx context.Context, cfg *config.Config) {

	db := database.DbConn(pctx, cfg).Database("booking_db")
	defer db.Client().Disconnect(pctx)

	col := db.Collection("booking_history")

	_, err := col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{"_id", 1}}},
		{Keys: bson.D{{"customer_id", 1}}},
	})

	if err != nil {
		log.Fatalf("Error: CreateIndex Failed")
		panic(err)
	}

	documents := func() []any {
		mockdatas := []*book.BookingHistory{
			{
				CustomerId: "customer0001",
				MovieName:  "Lalaland",
				Quantity:   1,
				Price:      150,
				BookingAt:  utils.GetLocaltime(),
				Seat: []book.Seat{
					{Number: "A3"},
				},
				ShowTime: "15.30",
			},
			{
				CustomerId: "customer0002",
				MovieName:  "Lalaland",
				Quantity:   1,
				Price:      150,
				BookingAt:  utils.GetLocaltime(),
				Seat: []book.Seat{
					{Number: "A4"},
				},
				ShowTime: "15.30",
			},
		}

		datas := make([]any, 0)

		for _, i := range mockdatas {
			datas = append(datas, i)
		}

		return datas

	}()

	result, err := col.InsertMany(pctx, documents)
	if err != nil {
		log.Printf("Insert BookingHistory failed:%s", err.Error())
		panic(err)
	}

	col = db.Collection("booking_transaction")

	_, err = col.InsertOne(pctx, bson.D{})

	if err != nil {
		log.Println("insert booking transaction failed")
		panic(err)
	}

	log.Println("Migrate Booking History completed:", result)

}
