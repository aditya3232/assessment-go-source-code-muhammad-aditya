package model

type CustomerResponse struct {
	ID            string `json:"id"`
	NationalId    int64  `json:"national_id"`
	Name          string `json:"name"`
	DetailAddress string `json:"detail_address"`
	CreatedAt     int64  `json:"created_at"`
	UpdatedAt     int64  `json:"updated_at"`
}

type CreateCustomerRequest struct {
	NationalId    int64  `json:"national_id" validate:"required"`
	Name          string `json:"name" validate:"required,max=255"`
	DetailAddress string `json:"detail_address" validate:"required,max=255"`
}

type UpdateCustomerRequest struct {
	ID            string `json:"-" validate:"required,max=100,uuid"`
	Name          string `json:"name" validate:"required,max=255"`
	DetailAddress string `json:"detail_address" validate:"required,max=255"`
}

type SearchCustomerRequest struct {
	NationalId int    `json:"national_id"`
	Name       string `json:"name" validate:"max=255"`
	Page       int    `json:"page" validate:"min=1"`
	Size       int    `json:"size" validate:"min=1,max=100"`
}

type GetCustomerRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}

type DeleteCustomerRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}
