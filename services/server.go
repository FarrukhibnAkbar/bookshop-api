package services

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"

	"github.com/FarrukhibnAkbar/bookshop-api/config"
	pb "github.com/FarrukhibnAkbar/bookshop-api/genproto/catalog_service"
	pb2 "github.com/FarrukhibnAkbar/bookshop-api/genproto/order_service"
)

type IServiceManager interface {
	CatalogService() pb.CatalogServiceClient
	OrderService() pb2.OrderServiceClient
}

type serviceManager struct {
	catalogService pb.CatalogServiceClient
	orderService pb2.OrderServiceClient
}

func (s *serviceManager) CatalogService() pb.CatalogServiceClient {
	return s.catalogService
}
func (s *serviceManager) OrderService() pb2.OrderServiceClient {
	return s.orderService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connCatalog, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.CatalogServiceHost, conf.CatalogServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connOrder, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.OrderServiceHost, conf.OrderServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		catalogService: pb.NewCatalogServiceClient(connCatalog),
		orderService: pb2.NewOrderServiceClient(connOrder),
	}

	return serviceManager, nil

}
