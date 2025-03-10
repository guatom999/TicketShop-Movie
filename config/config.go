package config

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App    App
		Db     Db
		AppUrl AppUrl
		Jwt    Jwt
		Kafka  Kafka
		Omise  Omise
		Gcp    Gcp
		Redis  Redis
		Mailer Mailer
	}

	App struct {
		Name string
		Port string
	}

	Db struct {
		Url string
	}
	AppUrl struct {
		CustomerUrl  string
		InventoryUrl string
		MovieUrl     string
		PaymentUrl   string
		TicketUrl    string
	}
	Jwt struct {
		AccessSecretKey  string
		RefreshSecretKey string
		ApiSecretKey     string
		AccessDuration   int64
		RefreshDuration  int64
	}

	Kafka struct {
		Url       string
		ApiKey    string
		SecretKey string
	}

	Omise struct {
		PublicKey string
		SecretKey string
	}

	Gcp struct {
		BucketName string
		FileLimit  int64
	}

	Redis struct {
		RedisUrl string
	}

	Mailer struct {
		MailerHost     string
		MailerPort     int
		MailerUserName string
		MailerPassword string
	}
)

func GetConfig(path string) Config {

	dir, file := filepath.Split(path)
	fmt.Println("dir and file is", dir, file)
	// fileName := strings.TrimPrefix(file, ".env.")

	viper.SetConfigName(file)
	viper.SetConfigType("env")
	viper.AddConfigPath(dir)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("fatal error config file: %s", err.Error())
		panic(err)
	}

	return Config{
		App: App{
			Name: viper.GetString("APP_NAME"),
			Port: viper.GetString("APP_PORT"),
		},
		Db: Db{
			Url: viper.GetString("DB_URL"),
		},
		AppUrl: AppUrl{
			CustomerUrl:  viper.GetString("TICKET_CUSTOMER_URL"),
			InventoryUrl: viper.GetString("TICKET_INVENTORY_URL"),
			MovieUrl:     viper.GetString("TICKET_MOVIE_URL"),
			PaymentUrl:   viper.GetString("TICKET_PAYMENT_URL"),
			TicketUrl:    viper.GetString("TICKET_TICKET_URL"),
		},
		Jwt: Jwt{
			AccessSecretKey:  viper.GetString("JWT_ACCESS_SECRET_KEY"),
			RefreshSecretKey: viper.GetString("JWT_REFRESH_SECRET_KEY"),
			ApiSecretKey:     viper.GetString("JWT_API_SECRET_KEY"),
			AccessDuration:   int64(viper.GetInt("JWT_ACCESS_DURATION")),
			RefreshDuration:  int64(viper.GetInt("JWT_REFRESH_DURATION")),
		},
		Kafka: Kafka{
			Url:       viper.GetString("KAFKA_URL"),
			ApiKey:    viper.GetString("KAFKA_API_KEY"),
			SecretKey: viper.GetString("KAFKA_SECRET_KEY"),
		},
		Omise: Omise{
			PublicKey: viper.GetString("OMISE_PUBLIC_KEY"),
			SecretKey: viper.GetString("OMISE_SECRET_KEY"),
		},
		Gcp: Gcp{
			BucketName: viper.GetString("APP_GCP_BUCKET"),
			FileLimit:  int64(viper.GetInt("APP_FILE_LIMIT")),
		},
		Redis: Redis{
			RedisUrl: viper.GetString("REDIS_URL"),
		},
		Mailer: Mailer{
			MailerHost:     viper.GetString("MAILER_HOST"),
			MailerPort:     viper.GetInt("MAILER_PORT"),
			MailerUserName: viper.GetString("MAILER_USERNAME"),
			MailerPassword: viper.GetString("MAILER_PASSWORD"),
		},
	}
}

func GetMigrateConfig(path string) Config {
	viper.SetConfigName(fmt.Sprintf(".env.%s", path))
	viper.SetConfigType("env")
	viper.AddConfigPath("../../env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("fatal error config file: %s", err.Error())
		panic(err)
	}

	return Config{
		App: App{
			Name: viper.GetString("APP_NAMe"),
			Port: viper.GetString("APP_PORT"),
		},
		Db: Db{
			Url: viper.GetString("DB_URL"),
		},
		Jwt: Jwt{
			AccessSecretKey:  viper.GetString("JWT_ACCESS_SECRET_KEY"),
			RefreshSecretKey: viper.GetString("JWT_REFRESH_SECRET_KEY"),
			ApiSecretKey:     viper.GetString("JWT_API_SECRET_KEY"),
			AccessDuration:   int64(viper.GetInt("JWT_ACCESS_DURATION")),
			RefreshDuration:  int64(viper.GetInt("JWT_REFRESH_DURATION")),
		},
		Kafka: Kafka{
			Url:       viper.GetString("KAFKA_URL"),
			ApiKey:    viper.GetString("KAFKA_API_KEY"),
			SecretKey: viper.GetString("KAFKA_SECRET_KEY"),
		},

		Omise: Omise{
			PublicKey: viper.GetString("OMISE_PUBLIC_KEY"),
			SecretKey: viper.GetString("OMISE_SECRET_KEY"),
		},
	}
}
