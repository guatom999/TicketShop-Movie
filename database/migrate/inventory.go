package migrate

import (
	"context"
	"fmt"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/database"
	"github.com/guatom999/TicketShop-Movie/modules/inventory"
)

func InventoryMigrate(pctx context.Context, cfg *config.Config) {

	db := database.DbConn(pctx, cfg).Database("inventory_db")
	defer db.Client().Disconnect(pctx)

	col := db.Collection("ticket_inventory")

	documents := func() []any {
		mockDatas := []inventory.Inventory{
			{
				CustomerId: "user0001",
				TicketId:   "Ticket0001",
			},
		}

		data := make([]any, 0)

		for _, v := range mockDatas {
			data = append(data, v)
		}

		return data

	}()

	if _, err := col.InsertMany(pctx, documents); err != nil {
		fmt.Println("Error: Insert Inventory failed :", err)
		panic(err)
	}

	fmt.Println("Migrate Inventory successfully")

}
