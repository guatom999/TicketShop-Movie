package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App   App
		Db    Db
		Jwt   Jwt
		Kafka Kafka
		Omise Omise
		Gcp   Gcp
	}

	App struct {
		Name string
		Port int
	}

	Db struct {
		Url string
	}

	Jwt struct {
		AccessSecretKey  string
		RefreshSecretKey string
		ApiSecretKey     string
		AccessDuration   int64
		RefreshDuration  int64
	}

	Kafka struct {
		Url   string
		Topic string
	}

	Omise struct {
		PublicKey string
		SecretKey string
	}

	Gcp struct {
		BucketName string
		FileLimit  int64
	}
)

func GetConfig(path string) Config {

	viper.SetConfigName(fmt.Sprintf(".env.%s", path))
	viper.SetConfigType("env")
	viper.AddConfigPath("./env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("fatal error config file: %s", err.Error())
		panic(err)
	}

	// fmt.Println("viper get env test ", viper.GetString("DB_HOST"))

	return Config{
		App: App{
			Name: viper.GetString("APP_NAMe"),
			Port: viper.GetInt("APP_PORT"),
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
			Url:   viper.GetString("KAFKA_URL"),
			Topic: viper.GetString("KAFKA_API_KEY"),
		},
		Omise: Omise{
			PublicKey: viper.GetString("OMISE_PUBLIC_KEY"),
			SecretKey: viper.GetString("OMISE_SECRET_KEY"),
		},
		Gcp: Gcp{
			BucketName: viper.GetString("APP_GCP_BUCKET"),
			FileLimit:  int64(viper.GetInt("APP_FILE_LIMIT")),
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

	// fmt.Println("viper get env test ", viper.GetString("DB_HOST"))

	return Config{
		App: App{
			Name: viper.GetString("APP_NAMe"),
			Port: viper.GetInt("APP_PORT"),
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
			Url:   viper.GetString("KAFKA_URL"),
			Topic: viper.GetString("KAFKA_API_KEY"),
		},

		Omise: Omise{
			PublicKey: viper.GetString("OMISE_PUBLIC_KEY"),
			SecretKey: viper.GetString("OMISE_SECRET_KEY"),
		},
	}
}
