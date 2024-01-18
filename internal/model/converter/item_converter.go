package converter

import (
	"assessment-go-source-code-muhammad-aditya/internal/entity"
	"assessment-go-source-code-muhammad-aditya/internal/model"
)

func ItemToResponse(item *entity.Item) *model.ItemResponse {
	return &model.ItemResponse{
		ID:        item.ID,
		ItemCode:  item.ItemCode,
		ItemName:  item.ItemName,
		Type:      item.Type,
		ItemPrice: item.ItemPrice,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}
