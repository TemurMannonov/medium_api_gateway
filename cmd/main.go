package main

import (
	"log"

	_ "github.com/lib/pq"

	"github.com/TemurMannonov/medium_api_gateway/api"
	"github.com/TemurMannonov/medium_api_gateway/config"
	grpcPkg "github.com/TemurMannonov/medium_api_gateway/pkg/grpc_client"
	"github.com/TemurMannonov/medium_api_gateway/pkg/logger"
)

func main() {
	cfg := config.Load(".")

	logrus := logger.New()

	grpcConn, err := grpcPkg.New(cfg)
	if err != nil {
		log.Fatalf("failed to get grpc connections: %v", err)
	}

	apiServer := api.New(&api.RouterOptions{
		Cfg:        &cfg,
		GrpcClient: grpcConn,
		Logger:     logrus,
	})

	err = apiServer.Run(cfg.HttpPort)
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
