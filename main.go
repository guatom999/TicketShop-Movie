package main

import (
	"context"
	"log"
	"os"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/database"
	"github.com/guatom999/TicketShop-Movie/database/redisConn"
	"github.com/guatom999/TicketShop-Movie/pkg/opn"
	"github.com/guatom999/TicketShop-Movie/server"
	"github.com/guatom999/TicketShop-Movie/utils"
)

func main() {

	ctx := context.Background()

	cfg := config.GetConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}
		log.Printf("choosen env is :%v", os.Args[1])
		return os.Args[1]
	}())

	db := database.DbConn(ctx, &cfg)

	defer db.Disconnect(ctx)

	redis := redisConn.RedisConn(ctx, &cfg)

	mailer := utils.ConnectToMailer(&cfg)

	omiseClient := opn.OmiseConn(&cfg)

	server.NewEchoServer(db, &cfg, omiseClient, redis, mailer).Start(ctx)

}
