package converter

import (
	"assessment-go-source-code-muhammad-aditya/internal/entity"
	"assessment-go-source-code-muhammad-aditya/internal/model"
)

func CustomerToResponse(customer *entity.Customer) *model.CustomerResponse {
	return &model.CustomerResponse{
		ID:            customer.ID,
		NationalId:    customer.NationalId,
		Name:          customer.Name,
		DetailAddress: customer.DetailAddress,
		CreatedAt:     customer.CreatedAt,
		UpdatedAt:     customer.UpdatedAt,
	}
}

func CustomerToEvent(customer *entity.Customer) *model.CustomerEvent {
	return &model.CustomerEvent{
		ID:            customer.ID,
		NationalId:    customer.NationalId,
		Name:          customer.Name,
		DetailAddress: customer.DetailAddress,
		CreatedAt:     customer.CreatedAt,
		UpdatedAt:     customer.UpdatedAt,
	}
}
