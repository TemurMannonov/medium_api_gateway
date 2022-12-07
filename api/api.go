package api

import (
	v1 "github.com/TemurMannonov/medium_api_gateway/api/v1"
	"github.com/TemurMannonov/medium_api_gateway/config"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "github.com/TemurMannonov/medium_api_gateway/api/docs" // for swagger

	grpcPkg "github.com/TemurMannonov/medium_api_gateway/pkg/grpc_client"
)

type RouterOptions struct {
	Cfg        *config.Config
	GrpcClient *grpcPkg.GrpcClient
}

// @title           Swagger for blog api
// @version         1.0
// @description     This is a blog service api.
// @host      localhost:8000
// @BasePath  /v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Security ApiKeyAuth
func New(opt *RouterOptions) *gin.Engine {
	router := gin.Default()

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Cfg:        opt.Cfg,
		GrpcClient: opt.GrpcClient,
	})

	apiV1 := router.Group("/v1")

	apiV1.POST("/users", handlerV1.CreateUser)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
