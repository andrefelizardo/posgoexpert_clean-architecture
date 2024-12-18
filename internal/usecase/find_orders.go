package usecase

import "github.com/andrefelizardo/posgoexpert_clean-architecture/internal/entity"

type FindOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewFindOrdersUseCase(OrderRepository entity.OrderRepositoryInterface) *FindOrdersUseCase {
	return &FindOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (l *FindOrdersUseCase) Execute() ([]OrderOutputDTO, error) {
	
	data, err := l.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var output []OrderOutputDTO

	for _, order := range data {
		dto := OrderOutputDTO{
			ID: order.ID,
			Price: order.Price,
			Tax: order.Tax,
			FinalPrice: order.FinalPrice,
		}
		output = append(output, dto)
	}
	
	return output, nil
}