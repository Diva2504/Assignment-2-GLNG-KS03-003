package models

import "gorm.io/gorm"

type Items struct {
	gorm.Model
	ItemCode    string `json:"items_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     uint   `json:"order_id"`
}
