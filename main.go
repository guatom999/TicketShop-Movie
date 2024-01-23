package main

import (
	"context"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/database"
	"github.com/guatom999/TicketShop-Movie/server"
)

func main() {

	ctx := context.Background()

	cfg := config.GetConfig()

	db := database.DbConn(ctx, &cfg)

	defer db.Disconnect(ctx)

	server.NewEchoServer(db, &cfg).Start(ctx)

	// 	func() string {
	// 	if len(os.Args) < 2 {
	// 		log.Fatal("Error: .env path is required")
	// 	}
	// 	log.Printf("choosen env is :%v", os.Args[1])
	// 	return os.Args[1]
	// }()

}
