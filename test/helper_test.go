package test

import (
	"assessment-go-source-code-muhammad-aditya/internal/entity"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func ClearAll() {
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

func CreateCustomers(t *testing.T, total int) {
	for i := 0; i < total; i++ {
		customer := &entity.Customer{
			ID:            uuid.NewString(),
			NationalId:    18071999,
			Name:          "Muhammad Aditya",
			DetailAddress: "jl.kebenaran jawa timur, malang",
		}
		err := db.Create(customer).Error
		assert.Nil(t, err)
	}
}

func CreateItems(t *testing.T, total int) {
	for i := 0; i < total; i++ {
		item := &entity.Item{
			ID:       uuid.NewString(),
			ItemCode: 1234,
			ItemName: "development",
			Type:     "service",
		}
		err := db.Create(item).Error
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
