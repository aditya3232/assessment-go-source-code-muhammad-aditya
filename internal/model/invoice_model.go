package model

type InvoiceResponse struct {
	ID            string                `json:"id"`
	InvoiceNumber string                `json:"invoice_number"`
	CustomerId    string                `json:"customer_id"`
	Subject       string                `json:"subject"`
	IssuedDate    int64                 `json:"issued_date"`
	DueDate       int64                 `json:"due_date"`
	SubTotal      int64                 `json:"sub_total"`
	GrandTotal    int64                 `json:"grand_total"`
	Status        string                `json:"status"`
	CreatedAt     int64                 `json:"created_at"`
	UpdatedAt     int64                 `json:"updated_at"`
	Customer      CustomerResponse      `json:"customer,omitempty"`
	InvoiceItems  []InvoiceItemResponse `json:"invoice_items,omitempty"`
}

type CreateInvoiceRequest struct {
	InvoiceNumber string                     `json:"invoice_number" validate:"required,max=255"`
	CustomerId    string                     `json:"customer_id" validate:"required,max=255"` //ambil dari parameter route
	Subject       string                     `json:"subject" validate:"required,max=255"`
	IssuedDate    int64                      `json:"issued_date" validate:"required"`
	DueDate       int64                      `json:"due_date" validate:"required"`
	SubTotal      int64                      `json:"sub_total" validate:"required"`
	GrandTotal    int64                      `json:"grand_total" validate:"required"`
	Status        string                     `json:"status" validate:"required"`
	InvoiceItems  []CreateInvoiceItemRequest `json:"invoice_items"`
}

type GetInvoiceRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}
