package converter

import (
	"assessment-go-source-code-muhammad-aditya/internal/entity"
	"assessment-go-source-code-muhammad-aditya/internal/model"
)

func CustomerToResponse(customer *entity.Customer) *model.CustomerResponse {
	return &model.CustomerResponse{
		ID:            customer.ID,
		Name:          customer.Name,
		DetailAddress: customer.DetailAddress,
		CreatedAt:     customer.CreatedAt,
		UpdatedAt:     customer.UpdatedAt,
	}
}
