package test

import (
	"assessment-go-source-code-muhammad-aditya/internal/entity"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func ClearAll() {
	ClearInvoiceItem()
	ClearInvoice()
	ClearCustomer()
	ClearItem()
}

func ClearCustomer() {
	err := db.Where("id is not null").Delete(&entity.Customer{}).Error
	if err != nil {
		log.Fatalf("Failed clear customer data : %+v", err)
	}
}

func ClearItem() {
	err := db.Where("id is not null").Delete(&entity.Item{}).Error
	if err != nil {
		log.Fatalf("Failed clear item data : %+v", err)
	}
}

func ClearInvoiceItem() {
	err := db.Where("id is not null").Delete(&entity.InvoiceItem{}).Error
	if err != nil {
		log.Fatalf("Failed clear invoice item data : %+v", err)
	}
}

func ClearInvoice() {
	err := db.Where("id is not null").Delete(&entity.Invoice{}).Error
	if err != nil {
		log.Fatalf("Failed clear invoice data : %+v", err)
	}
}

func CreateCustomers(t *testing.T, total int) {
	for i := 0; i < total; i++ {
		customer := &entity.Customer{
			ID:            uuid.NewString(),
			NationalId:    int64(18071999 + i),
			Name:          fmt.Sprintf("Muhammad Aditya %d", i),
			DetailAddress: fmt.Sprintf("jl.kebenaran jawa timur, malang no: % d", i),
		}
		err := db.Create(customer).Error
		assert.Nil(t, err)
	}
}

func CreateItems(t *testing.T, total int) {
	for i := 0; i < total; i++ {
		item := &entity.Item{
			ID:       uuid.NewString(),
			ItemCode: int64(1234 + i),
			ItemName: fmt.Sprintf("development %d", i),
			Type:     fmt.Sprintf("service %d", i),
		}
		err := db.Create(item).Error
		assert.Nil(t, err)
	}
}

func CreateInvoices(t *testing.T, customer *entity.Customer, total int) {
	for i := 0; i < total; i++ {
		invoice := &entity.Invoice{
			ID:            uuid.NewString(),
			InvoiceNumber: fmt.Sprintf("0002 %d", i),
			CustomerId:    customer.ID,
			Subject:       fmt.Sprintf("winter marketing campaign %d", i),
			IssuedDate:    int64(170556121558 + i),
			DueDate:       int64(170556053649 + i),
			TotalItem:     int64(3),
			SubTotal:      int64(1400),
			GrandTotal:    int64(1540),
			Status:        "paid",
		}
		err := db.Create(invoice).Error
		assert.Nil(t, err)
	}
}

func CreateInvoiceItems(t *testing.T, invoice *entity.Invoice, item *entity.Item, total int) {
	for i := 0; i < total; i++ {
		invoiceItem := &entity.InvoiceItem{
			ID:           uuid.NewString(),
			InvoiceId:    invoice.ID,
			ItemId:       item.ID,
			ItemPrice:    int64(500 + i),
			ItemQuantity: int64(1 + i),
			Amount:       int64(500 + i),
		}
		err := db.Create(invoiceItem).Error
		assert.Nil(t, err)
	}
}

func GetFirstCustomer(t *testing.T) *entity.Customer {
	customer := new(entity.Customer)
	err := db.First(customer).Error
	assert.Nil(t, err)
	return customer
}

func GetFirstItem(t *testing.T) *entity.Item {
	item := new(entity.Item)
	err := db.First(item).Error
	assert.Nil(t, err)
	return item
}

func GetFirstInvoice(t *testing.T) *entity.Invoice {
	invoice := new(entity.Invoice)
	err := db.First(invoice).Error
	assert.Nil(t, err)
	return invoice
}

func GetAllItem(t *testing.T) []*entity.Item {
	items := make([]*entity.Item, 0)
	err := db.Find(&items).Error
	assert.Nil(t, err)
	return items
}
