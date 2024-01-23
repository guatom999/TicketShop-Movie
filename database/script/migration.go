package main

import (
	"context"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/database/migrate"
)

func main() {
	ctx := context.Background()

	cfg := config.GetConfig()

	migrate.MovieMigrate(ctx, &cfg)

}
