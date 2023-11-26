package model

import "time"

type Item struct {
	ItemID      string    `gorm:"primaryKey" json:"itemId"`
	ItemCode    string    `json:"itemCode"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	OrderID     int       `json:"orderID" gorm:"foreignKey:OrderID"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
