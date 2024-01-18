package repository

import (
	"assessment-go-source-code-muhammad-aditya/internal/entity"
	"assessment-go-source-code-muhammad-aditya/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ItemRepository struct {
	Repository[entity.Item]
	Log *logrus.Logger
}

func NewItemRepository(log *logrus.Logger) *ItemRepository {
	return &ItemRepository{
		Log: log,
	}
}

func (r *ItemRepository) CountByItemCode(db *gorm.DB, item *entity.Item) (int64, error) {
	var total int64
	err := db.Model(item).Where("item_code = ?", item.ItemCode).Count(&total).Error
	return total, err
}

func (r *ItemRepository) Search(db *gorm.DB, request *model.SearchItemRequest) ([]entity.Item, int64, error) {
	var items []entity.Item
	if err := db.Scopes(r.FilterItem(request)).Offset((request.Page - 1) * request.Size).Limit(request.Size).Find(&items).Error; err != nil {
		return nil, 0, err
	}

	var total int64 = 0
	if err := db.Model(&entity.Item{}).Scopes(r.FilterItem(request)).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

func (r *ItemRepository) FilterItem(request *model.SearchItemRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if item_code := request.ItemCode; item_code != 0 {
			tx = tx.Where("item_code = ?", item_code)
		}

		if item_name := request.ItemName; item_name != "" {
			item_name = "%" + item_name + "%"
			tx = tx.Where("item_name LIKE ?", item_name)
		}

		if item_price := request.ItemPrice; item_price != 0 {
			tx = tx.Where("item_price = ?", item_price)
		}

		if item_type := request.Type; item_type != "" {
			item_type = "%" + item_type + "%"
			tx = tx.Where("type LIKE ?", item_type)
		}

		return tx
	}
}
