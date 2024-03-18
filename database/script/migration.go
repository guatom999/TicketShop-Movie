package main

import (
	"context"
	"log"
	"os"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/database/migrate"
)

type Config struct {
	Db Db
}

type Db struct {
	Url string
}

func main() {
	ctx := context.Background()

	path := func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}
		log.Printf("choosen env is :%v", os.Args[1])
		return os.Args[1]
	}()

	cfg := config.GetMigrateConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}
		log.Printf("choosen env is :%v", os.Args[1])
		return os.Args[1]
	}())

	switch path {
	case "movie":
		migrate.MovieMigrate(ctx, &cfg)
	case "booking":
		migrate.BookingMigrate(ctx, &cfg)
	case "inventory":
		migrate.InventoryMigrate(ctx, &cfg)
	case "customer":
		migrate.CustomerMigrate(ctx, &cfg)
	case "ticket":
		migrate.TicketMigrate(ctx, &cfg)
	case "payment":
		migrate.PaymentMigrate(ctx, &cfg)
	}

}

// func GetConfig(path string) config.Config {
// 	viper.SetConfigName(fmt.Sprintf(".env.%s", path))
// 	viper.SetConfigType("env")

// 	viper.AddConfigPath("../../env")
// 	viper.AutomaticEnv()

// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		log.Fatalf("fatal error config file: %s", err.Error())
// 		panic(err)
// 	}

// 	return Config{
// 		Db: Db{
// 			Url: viper.GetString("DB_URL"),
// 		},
// 	}

// }
