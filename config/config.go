package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	HttpPort            string
	UserServiceGrpcPort string
	UserServiceHost     string
}

func Load(path string) Config {
	godotenv.Load(path + "/.env") // load .env file if it exists

	conf := viper.New()
	conf.AutomaticEnv()

	cfg := Config{
		HttpPort:            conf.GetString("HTTP_PORT"),
		UserServiceHost:     conf.GetString("USER_SERVICE_HOST"),
		UserServiceGrpcPort: conf.GetString("USER_SERVICE_GRPC_PORT"),
	}

	return cfg
}
