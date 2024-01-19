package repository

import (
	"assessment-go-source-code-muhammad-aditya/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type InvoiceItemRepository struct {
	Repository[entity.InvoiceItem]
	Log *logrus.Logger
}

func NewInvoiceItemRepository(log *logrus.Logger) *InvoiceItemRepository {
	return &InvoiceItemRepository{
		Log: log,
	}
}

// find by invoice_id
func (c *InvoiceItemRepository) FindByInvoiceId(tx *gorm.DB, invoiceItem *[]entity.InvoiceItem, invoiceId string) error {
	if err := tx.Where("invoice_id = ?", invoiceId).Find(invoiceItem).Error; err != nil {
		return err
	}

	return nil
}
