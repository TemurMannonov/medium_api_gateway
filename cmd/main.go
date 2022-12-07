package main

import (
	"log"

	_ "github.com/lib/pq"

	"github.com/TemurMannonov/blog/api"
	"github.com/TemurMannonov/blog/config"
)

func main() {
	cfg := config.Load(".")

	apiServer := api.New(&api.RouterOptions{
		Cfg: &cfg,
	})

	err := apiServer.Run(cfg.HttpPort)
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
