package migrate

import (
	"context"
	"log"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/database"
	"github.com/guatom999/TicketShop-Movie/modules/customer"
)

func CustomerMigrate(pctx context.Context, cfg *config.Config) {
	db := database.DbConn(pctx, cfg).Database("customer_db")
	defer db.Client().Disconnect(pctx)

	col := db.Collection("customer")

	documents := func() []any {
		mocksDatas := []customer.Customer{
			{
				Email:    "test1234@hotamil.com",
				UserName: "customer1",
			},
		}

		datas := make([]any, 0)

		for _, v := range mocksDatas {
			datas = append(datas, v)

		}

		return datas
	}()

	result, err := col.InsertMany(pctx, documents)
	if err != nil {
		log.Printf("Insert Customer failed:%s", err.Error())
		panic(err)
	}

	log.Println("BookingHistoryMigrate Customer Db completed", result)
}
