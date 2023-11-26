package model

import "time"

type Order struct {
	OrderID      int       `json:"orderID" gorm:"primaryKey"`
	OrderedAt    time.Time `json:"orderedAt"`
	CustomerName string    `json:"customerName"`
	Items        []Item    `json:"items" gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
