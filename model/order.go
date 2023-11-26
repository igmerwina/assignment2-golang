package model

import "time"

type Order struct {
	OrderID      int       `json:"OrderID" gorm:"primaryKey"`
	OrderedAt    time.Time `json:"orderedAt"`
	CustomerName string    `json:"customerName"`
	Items        []Item    `json:"items" gorm:"foreignKey:OrderID" `
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
