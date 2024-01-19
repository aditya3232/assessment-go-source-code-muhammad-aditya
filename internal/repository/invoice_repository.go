package repository

import (
	"assessment-go-source-code-muhammad-aditya/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type InvoiceRepository struct {
	Repository[entity.Invoice]
	Log *logrus.Logger
}

func NewInvoiceRepository(log *logrus.Logger) *InvoiceRepository {
	return &InvoiceRepository{
		Log: log,
	}
}

func (r *InvoiceRepository) CountByInvoiceNumber(db *gorm.DB, invoice *entity.Invoice) (int64, error) {
	var total int64
	err := db.Model(invoice).Where("invoice_number = ?", invoice.InvoiceNumber).Count(&total).Error
	return total, err
}
