package v1

import (
	"github.com/TemurMannonov/blog/api/models"
	"github.com/TemurMannonov/blog/config"
)

type handlerV1 struct {
	cfg *config.Config
}

type HandlerV1Options struct {
	Cfg *config.Config
}

func New(options *HandlerV1Options) *handlerV1 {
	return &handlerV1{
		cfg: options.Cfg,
	}
}

func errorResponse(err error) *models.ErrorResponse {
	return &models.ErrorResponse{
		Error: err.Error(),
	}
}
