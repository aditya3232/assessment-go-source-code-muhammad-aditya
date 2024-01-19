package model

type InvoiceItemResponse struct {
	ID           string `json:"id"`
	InvoiceId    string `json:"invoice_id"`
	ItemId       string `json:"item_id"`
	ItemPrice    int64  `json:"item_price"`
	ItemQuantity int64  `json:"item_quantity"`
	Amount       int64  `json:"amount"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
}

type CreateInvoiceItemRequest struct {
	InvoiceId    string `json:"-" validate:"required,max=255"`
	ItemId       string `json:"item_id" validate:"required,max=255"`
	ItemPrice    int64  `json:"item_price" validate:"required"`
	ItemQuantity int64  `json:"item_quantity" validate:"required"`
	Amount       int64  `json:"amount" validate:"required"`
}

type GetInvoiceItemRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}

type DeleteInvoiceItemRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}
