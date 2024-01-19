package converter

import (
	"assessment-go-source-code-muhammad-aditya/internal/entity"
	"assessment-go-source-code-muhammad-aditya/internal/model"
)

func InvoiceItemToResponse(invoiceitem *entity.InvoiceItem) *model.InvoiceItemResponse {
	return &model.InvoiceItemResponse{
		ID:           invoiceitem.ID,
		InvoiceId:    invoiceitem.InvoiceId,
		ItemId:       invoiceitem.ItemId,
		ItemPrice:    invoiceitem.ItemPrice,
		ItemQuantity: invoiceitem.ItemQuantity,
		Amount:       invoiceitem.Amount,
		CreatedAt:    invoiceitem.CreatedAt,
		UpdatedAt:    invoiceitem.UpdatedAt,
	}
}
