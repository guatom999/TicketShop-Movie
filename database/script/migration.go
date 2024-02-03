package main

import (
	"context"
	"log"
	"os"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/database/migrate"
)

func main() {
	ctx := context.Background()

	path := func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}
		log.Printf("choosen env is :%v", os.Args[1])
		return os.Args[1]
	}()

	cfg := config.GetConfig(path)

	switch path {
	case "movie":
		migrate.MovieMigrate(ctx, &cfg)
	case "booking":
		migrate.BookingMigrate(ctx, &cfg)
	}

}
