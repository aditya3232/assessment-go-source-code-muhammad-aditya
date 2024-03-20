package repository

import (
	"assessment-go-source-code-muhammad-aditya/internal/entity"
	"assessment-go-source-code-muhammad-aditya/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	Repository[entity.Customer] // embed generic repository, so we can use all generic method
	Log                         *logrus.Logger
}

func NewCustomerRepository(log *logrus.Logger) *CustomerRepository {
	return &CustomerRepository{
		Log: log,
	}
}

func (r *CustomerRepository) CountByNationalId(db *gorm.DB, customer *entity.Customer) (int64, error) {
	var total int64
	err := db.Model(customer).Where("national_id = ?", customer.NationalId).Count(&total).Error
	return total, err
}

func (r *CustomerRepository) Search(db *gorm.DB, request *model.SearchCustomerRequest) ([]entity.Customer, int64, error) {
	var customers []entity.Customer
	if err := db.Scopes(r.FilterCustomer(request)).Offset((request.Page - 1) * request.Size).Limit(request.Size).Find(&customers).Error; err != nil {
		return nil, 0, err
	}

	var total int64 = 0
	if err := db.Model(&entity.Customer{}).Scopes(r.FilterCustomer(request)).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return customers, total, nil
}

func (r *CustomerRepository) FilterCustomer(request *model.SearchCustomerRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if national_id := request.NationalId; national_id != 0 {
			tx = tx.Where("national_id = ?", national_id)
		}

		if name := request.Name; name != "" {
			name = "%" + name + "%"
			tx = tx.Where("name LIKE ?", name)
		}

		return tx
	}
}
