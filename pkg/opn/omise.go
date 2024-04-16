package opn

import (
	"log"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/omise/omise-go"
)

func OmiseConn(cfg *config.Config) *omise.Client {

	conn, err := omise.NewClient(cfg.Omise.PublicKey, cfg.Omise.SecretKey)
	if err != nil {
		log.Fatalf("failed to connect omise")
		panic(err)
	}

	return conn
}
