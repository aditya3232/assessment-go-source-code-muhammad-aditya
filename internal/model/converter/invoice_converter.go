package converter

import (
	"assessment-go-source-code-muhammad-aditya/internal/entity"
	"assessment-go-source-code-muhammad-aditya/internal/model"
	"time"
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
		TotalItem:     invoice.TotalItem,
		SubTotal:      invoice.SubTotal,
		GrandTotal:    invoice.GrandTotal,
		Status:        invoice.Status,
		CreatedAt:     invoice.CreatedAt,
		UpdatedAt:     invoice.UpdatedAt,
		Customer:      model.CustomerResponse(invoice.Customer),
		InvoiceItems:  convertedItems,
	}
}

func InvoiceListToResponse(invoice *entity.Invoice) *model.InvoiceListResponse {
	issuedDate := time.Unix(0, invoice.IssuedDate*int64(time.Millisecond)).Format("01/02/2006")
	dueDate := time.Unix(0, invoice.DueDate*int64(time.Millisecond)).Format("01/02/2006")

	return &model.InvoiceListResponse{
		InvoiceNumber: invoice.InvoiceNumber,
		IssuedDate:    issuedDate,
		Subject:       invoice.Subject,
		TotalItem:     invoice.TotalItem,
		CustomerName:  invoice.Customer.Name,
		DueDate:       dueDate,
		Status:        invoice.Status,
	}
}
