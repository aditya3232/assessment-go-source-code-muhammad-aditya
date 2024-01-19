package entity

type InvoiceItem struct {
	ID           string `gorm:"column:id;primaryKey"`
	InvoiceId    string `gorm:"column:invoice_id"`
	ItemId       string `gorm:"column:item_id"`
	ItemPrice    int64  `gorm:"column:item_price"`
	ItemQuantity int64  `gorm:"column:item_quantity"`
	Amount       int64  `gorm:"column:amount"`
	CreatedAt    int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt    int64  `gorm:"column:updated_at;autoUpdateTime:milli"`
}

func (i *InvoiceItem) TableName() string {
	return "invoice_items"
}
