package migrate

import (
	"context"
	"fmt"
	"log"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/database"
	"github.com/guatom999/TicketShop-Movie/modules/ticket"
	"github.com/guatom999/TicketShop-Movie/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TicketMigrate(pctx context.Context, cfg *config.Config) {

	db := database.DbConn(pctx, cfg).Database("ticket_db")
	defer db.Client().Disconnect(pctx)

	col := db.Collection("customer_ticket")

	_, err := col.Indexes().CreateOne(pctx, mongo.IndexModel{
		Keys: bson.D{{"_id", 1}},
	})
	if err != nil {
		fmt.Println("Error : CreateIndex Failed", err)
		panic(err)
	}

	documents := func() []any {
		mockdatas := []*ticket.Ticket{
			{
				MovieId:    "65ecc8b289430838d51441a2",
				CustomerId: "65e8a18968027287072e87dd",
				Seat:       "A1",
				MovieName:  "FRIEREN",
				Date:       "14/3/2024",
				Time:       "15.30",
				Price:      160,
				CreatedAt:  utils.GetLocaltime(),
				UpdatedAt:  utils.GetLocaltime(),
			},
		}

		datas := make([]any, 0)
		for i := range mockdatas {
			datas = append(datas, mockdatas[i])
		}

		return datas

	}()

	results, err := col.InsertMany(pctx, documents)
	if err != nil {
		log.Fatalf("Error: CreateIndex Failed %s", err.Error())
		panic(err)
	}

	log.Println("Migrate movies completed:", results)

}
