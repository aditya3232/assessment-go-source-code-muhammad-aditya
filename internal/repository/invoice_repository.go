package repository

import (
	"assessment-go-source-code-muhammad-aditya/internal/entity"
	"assessment-go-source-code-muhammad-aditya/internal/model"

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

func (r *InvoiceRepository) Search(db *gorm.DB, request *model.SearchInvoiceRequest) ([]entity.Invoice, int64, error) {
	var invoices []entity.Invoice
	if err := db.Scopes(r.FilterInvoice(request)).Offset((request.Page - 1) * request.Size).Limit(request.Size).Find(&invoices).Error; err != nil {
		return nil, 0, err
	}

	var total int64 = 0
	if err := db.Model(&entity.Invoice{}).Scopes(r.FilterInvoice(request)).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return invoices, total, nil
}

func (r *InvoiceRepository) FilterInvoice(request *model.SearchInvoiceRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if invoice_number := request.InvoiceNumber; invoice_number != "" {
			tx = tx.Where("invoice_number = ?", invoice_number)
		}

		if issuedDate := request.IssuedDate; issuedDate != 0 {
			tx = tx.Where("issued_date >= ?", issuedDate).Where("issued_date <= ?", issuedDate+86400000) // add one day
		}

		if dueDate := request.DueDate; dueDate != 0 {
			tx = tx.Where("due_date >= ?", dueDate).Where("due_date <= ?", dueDate+86400000)
		}

		if subject := request.Subject; subject != "" {
			subject = "%" + subject + "%"
			tx = tx.Where("subject LIKE ?", subject)
		}

		if total_item := request.TotalItem; total_item != 0 {
			tx = tx.Where("total_item = ?", total_item)
		}

		if customer_id := request.CustomerId; customer_id != "" {
			tx = tx.Where("customer_id = ?", customer_id)
		}

		if status := request.Status; status != "" {
			status = "%" + status + "%"
			tx = tx.Where("status LIKE ?", status)
		}

		return tx
	}
}
