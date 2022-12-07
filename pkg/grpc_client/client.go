package grpc_client

// import (
// 	"fmt"
// 	pbu "genproto/user_service"

// 	"google.golang.org/grpc"

// 	"/config"
// )

// type GrpcClientI interface {
// 	UserService() pbu.UserServiceClient
// }

// // GrpcClient ...
// type GrpcClient struct {
// 	cfg         config.Config
// 	connections map[string]interface{}
// }

// // New ...
// func New(cfg config.Config) (*GrpcClient, error) {

// 	connUser, err := grpc.Dial(
// 		fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort),
// 		grpc.WithInsecure(),
// 	)
// 	if err != nil {
// 		return nil, fmt.Errorf("user service dial host: %s port:%d err: %s",
// 			cfg.UserServiceHost, cfg.UserServicePort, err)
// 	}

// 	return &GrpcClient{
// 		cfg: cfg,
// 		connections: map[string]interface{}{
// 			"customer_service": pbu.NewCustomerServiceClient(connUser),
// 		},
// 	}, nil
// }

// // CustomerService ...
// func (g *GrpcClient) CustomerService() pbu.CustomerServiceClient {
// 	return g.connections["customer_service"].(pbu.CustomerServiceClient)
// }
