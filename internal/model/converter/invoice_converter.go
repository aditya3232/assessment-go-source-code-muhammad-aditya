package converter

import (
	"assessment-go-source-code-muhammad-aditya/internal/entity"
	"assessment-go-source-code-muhammad-aditya/internal/model"
)

func InvoiceToResponse(invoice *entity.Invoice) *model.InvoiceResponse {
	convertedItems := make([]model.InvoiceItemResponse, len(invoice.InvoiceItems))
	for i, item := range invoice.InvoiceItems {
		convertedItems[i] = model.InvoiceItemResponse(item)
	}
	return &model.InvoiceResponse{
		ID:            invoice.ID,
		InvoiceNumber: invoice.InvoiceNumber,
		CustomerId:    invoice.CustomerId,
		Subject:       invoice.Subject,
		IssuedDate:    invoice.IssuedDate,
		DueDate:       invoice.DueDate,
		SubTotal:      invoice.SubTotal,
		GrandTotal:    invoice.GrandTotal,
		Status:        invoice.Status,
		CreatedAt:     invoice.CreatedAt,
		UpdatedAt:     invoice.UpdatedAt,
		Customer:      model.CustomerResponse(invoice.Customer),
		InvoiceItems:  convertedItems,
	}
}
