package config

import (
	"log"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App App
		Db  Db
		Jwt Jwt
	}

	App struct {
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
)

func GetConfig() Config {

	viper.SetConfigName(".env.movie")
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
	}
}
