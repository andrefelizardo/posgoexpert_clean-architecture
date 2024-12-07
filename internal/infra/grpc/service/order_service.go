package service

import (
	"context"

	"github.com/andrefelizardo/posgoexpert_clean-architecture/internal/infra/grpc/pb"
	"github.com/andrefelizardo/posgoexpert_clean-architecture/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID: in.Id,
		Price: float64(in.Price),
		Tax: float64(in.Tax),
	}

	outupt, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}

	return &pb.CreateOrderResponse{
		Id: outupt.ID,
		Price: float32(outupt.Price),
		Tax: float32(outupt.Tax),
		FinalPrice: float32(outupt.FinalPrice),
	}, nil
}