package controllers

import (
	"gorm.io/gorm"
)

type Handler struct {
	Connect *gorm.DB
}
