package entity

type Invoice struct {
	ID            string        `gorm:"column:id;primaryKey"`
	InvoiceNumber string        `gorm:"column:invoice_number"`
	CustomerId    string        `gorm:"column:customer_id"`
	Subject       string        `gorm:"column:subject"`
	IssuedDate    int64         `gorm:"column:issued_date"`
	DueDate       int64         `gorm:"column:due_date"`
	TotalItem     int64         `gorm:"column:total_item"`
	SubTotal      int64         `gorm:"column:sub_total"`
	GrandTotal    int64         `gorm:"column:grand_total"`
	Status        string        `gorm:"column:status"`
	CreatedAt     int64         `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt     int64         `gorm:"column:updated_at;autoUpdateTime:milli"`
	Customer      Customer      `gorm:"foreignKey:customer_id;references:id"`
	InvoiceItems  []InvoiceItem `gorm:"foreignKey:invoice_id;references:id"`
}

func (i *Invoice) TableName() string {
	return "invoices"
}
