package model

type CustomerEvent struct {
	ID            string `json:"id"`
	NationalId    int64  `json:"national_id"`
	Name          string `json:"name"`
	DetailAddress string `json:"detail_address"`
	CreatedAt     int64  `json:"created_at"`
	UpdatedAt     int64  `json:"updated_at"`
}

func (c *CustomerEvent) GetId() string {
	return c.ID
}
