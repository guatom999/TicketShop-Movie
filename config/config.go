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
	}
}
