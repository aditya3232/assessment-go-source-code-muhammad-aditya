package entity

type Item struct {
	ID        string `gorm:"column:id;primaryKey"`
	ItemCode  int64  `gorm:"column:item_code"`
	ItemName  string `gorm:"column:item_name"`
	Type      string `gorm:"column:type"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime:milli"`
}

func (i *Item) TableName() string {
	return "items"
}
