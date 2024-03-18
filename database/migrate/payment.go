package migrate

import (
	"context"
	"log"
	"time"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/database"
	"go.mongodb.org/mongo-driver/bson"
)

func PaymentMigrate(pctx context.Context, cfg *config.Config) {

	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	db := database.DbConn(pctx, cfg).Database("payment_db")
	defer db.Client().Disconnect(pctx)

	col := db.Collection("payment_queue")
	result, err := col.InsertOne(ctx, bson.M{"offset": -1})

	if err != nil {
		panic(err)
	}

	log.Println("Migrate Payment Db successfully", result)

}
