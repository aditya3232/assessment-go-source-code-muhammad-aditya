package model

type ItemResponse struct {
	ID        string `json:"id"`
	ItemName  string `json:"item_name"`
	Type      string `json:"type"`
	ItemPrice int64  `json:"item_price"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type CreateItemRequest struct {
	ItemName  string `json:"item_name" validate:"required,max=255"`
	Type      string `json:"type" validate:"required,max=255"`
	ItemPrice int64  `json:"item_price" validate:"required"`
}

type UpdateItemRequest struct {
	ID        string `json:"-" validate:"required,max=100,uuid"`
	ItemName  string `json:"item_name" validate:"max=255"`
	Type      string `json:"type" validate:"max=255"`
	ItemPrice int64  `json:"item_price"`
}

type SearchItemRequest struct {
	ItemName  string `json:"item_name" validate:"max=255"`
	Type      string `json:"type" validate:"max=255"`
	ItemPrice int    `json:"item_price"`
	Page      int    `json:"page" validate:"min=1"`
	Size      int    `json:"size" validate:"min=1,max=100"`
}

type GetItemRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}

type DeleteItemRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}
