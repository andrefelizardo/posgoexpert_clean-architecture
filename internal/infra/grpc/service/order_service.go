package service

import (
	"context"

	"github.com/andrefelizardo/posgoexpert_clean-architecture/internal/infra/grpc/pb"
	"github.com/andrefelizardo/posgoexpert_clean-architecture/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	FindOrdersUseCase usecase.FindOrdersUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, findOrdersUseCase usecase.FindOrdersUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		FindOrdersUseCase: findOrdersUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID: in.Id,
		Price: float64(in.Price),
		Tax: float64(in.Tax),
	}

	outupt, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}

	return &pb.OrderResponse{
		Id: outupt.ID,
		Price: float32(outupt.Price),
		Tax: float32(outupt.Tax),
		FinalPrice: float32(outupt.FinalPrice),
	}, nil
}

func (s *OrderService) FindAllOrders(ctx context.Context, in *pb.Empty) (*pb.ListOrders, error) {
	output, err := s.FindOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var orders []*pb.OrderResponse

	for _, item := range output {
		order := &pb.OrderResponse{
			Id: item.ID,
			Price: float32(item.Price),
			Tax: float32(item.Tax),
			FinalPrice: float32(item.FinalPrice),
		}
		orders = append(orders, order)
	}

	return &pb.ListOrders{
		Orders: orders,
	}, nil
}