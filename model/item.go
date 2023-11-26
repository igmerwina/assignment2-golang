package model

import "time"

type Item struct {
	ItemID      string    `gorm:"primaryKey" json:"itemId"`
	ItemCode    string    `json:"itemCode"`
	OrderID     int       `json:"orderId"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
