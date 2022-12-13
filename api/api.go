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
	GrpcClient grpcPkg.GrpcClientI
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

	apiV1.POST("/auth/register", handlerV1.Register)
	apiV1.POST("/auth/verify", handlerV1.Verify)

	apiV1.POST("/users", handlerV1.CreateUser)
	apiV1.GET("/users/:id", handlerV1.GetUser)
	apiV1.PUT("/users/:id", handlerV1.UpdateUser)
	apiV1.GET("/users", handlerV1.GetAllUsers)
	apiV1.DELETE("/users/:id", handlerV1.DeleteUser)
	apiV1.GET("/users/email/:email", handlerV1.GetUserByEmail)

	apiV1.POST("/posts", handlerV1.AuthMiddleware, handlerV1.CreatePost)

	apiV1.POST("/categories", handlerV1.CreateCategory)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
