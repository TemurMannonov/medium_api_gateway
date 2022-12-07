package v1

import (
	"github.com/TemurMannonov/medium_api_gateway/api/models"
	"github.com/TemurMannonov/medium_api_gateway/config"
	grpcPkg "github.com/TemurMannonov/medium_api_gateway/pkg/grpc_client"
)

type handlerV1 struct {
	cfg        *config.Config
	grpcClient *grpcPkg.GrpcClient
}

type HandlerV1Options struct {
	Cfg        *config.Config
	GrpcClient *grpcPkg.GrpcClient
}

func New(options *HandlerV1Options) *handlerV1 {
	return &handlerV1{
		cfg:        options.Cfg,
		grpcClient: options.GrpcClient,
	}
}

func errorResponse(err error) *models.ErrorResponse {
	return &models.ErrorResponse{
		Error: err.Error(),
	}
}
