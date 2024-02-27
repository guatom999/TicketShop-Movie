package migrate

import (
	"context"

	"github.com/guatom999/TicketShop-Movie/config"
)

func CustomerMigrate(pctx context.Context, cfg *config.Config) {
	// db := database.DbConn(pctx, cfg).Database("movie_db")
	// defer db.Client().Disconnect(pctx)

	// col := db.Collection("customer")

	// documents := func() []any {

	// 		mockDatas := []

	// 	}()
}
