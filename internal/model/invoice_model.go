package model

type InvoiceResponse struct {
	ID            string                `json:"id"`
	InvoiceNumber string                `json:"invoice_number"`
	CustomerId    string                `json:"customer_id"`
	Subject       string                `json:"subject"`
	IssuedDate    int64                 `json:"issued_date"`
	DueDate       int64                 `json:"due_date"`
	TotalItem     int64                 `json:"total_item"`
	SubTotal      int64                 `json:"sub_total"`
	GrandTotal    int64                 `json:"grand_total"`
	Status        string                `json:"status"`
	CreatedAt     int64                 `json:"created_at"`
	UpdatedAt     int64                 `json:"updated_at"`
	Customer      CustomerResponse      `json:"customer,omitempty"`
	InvoiceItems  []InvoiceItemResponse `json:"invoice_items,omitempty"`
}

type InvoiceListResponse struct {
	InvoiceNumber string `json:"invoice_id"`
	IssuedDate    string `json:"issued_date"`
	Subject       string `json:"subject"`
	TotalItem     int64  `json:"total_item"`
	CustomerName  string `json:"customer"`
	DueDate       string `json:"due_date"`
	Status        string `json:"status"`
}

type CreateInvoiceRequest struct {
	InvoiceNumber string                     `json:"invoice_number" validate:"required,max=255"`
	CustomerId    string                     `json:"customer_id" validate:"required,max=255"` //ambil dari parameter route
	Subject       string                     `json:"subject" validate:"required,max=255"`
	IssuedDate    int64                      `json:"issued_date" validate:"required"`
	DueDate       int64                      `json:"due_date" validate:"required"`
	TotalItem     int64                      `json:"total_item" validate:"required"`
	SubTotal      int64                      `json:"sub_total" validate:"required"`
	GrandTotal    int64                      `json:"grand_total" validate:"required"`
	Status        string                     `json:"status" validate:"required"`
	InvoiceItems  []CreateInvoiceItemRequest `json:"invoice_items"`
}

type SearchInvoiceRequest struct {
	InvoiceNumber string `json:"invoice_number" validate:"max=255"`
	IssuedDate    int64  `json:"issued_date"`
	DueDate       int64  `json:"due_date"`
	Subject       string `json:"subject" validate:"max=255"`
	TotalItem     int    `json:"total_item"`
	CustomerId    string `json:"customer_id" validate:"max=255"`
	Status        string `json:"status" validate:"max=255"`
	Page          int    `json:"page" validate:"min=1"`
	Size          int    `json:"size" validate:"min=1,max=100"`
}

type GetInvoiceRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}
