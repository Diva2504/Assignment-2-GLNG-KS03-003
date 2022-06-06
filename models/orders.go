package models

import (
	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	Customer_name string
	Items         *Items `gorm:"Foreignkey:OrderID;association_foreignkey:ID;"`
}
