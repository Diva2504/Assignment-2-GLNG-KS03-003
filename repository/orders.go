package repository

import (
	"github.com/Diva2504/Assignment-2-GLNG-KS03-003/models"
	"gorm.io/gorm"
)

func GetAllOrders(db *gorm.DB) ([]models.Orders, error) {
	var orders []models.Orders
	result := db.Find(&orders)

	if result.Error != nil {
		return nil, result.Error
	} else {
		if result.RowsAffected <= 0 {
			return nil, result.Error
		} else {
			return orders, result.Error
		}
	}

}

func CreateOrder(input *models.Orders, db *gorm.DB) error {
	result := db.Debug().Create(&input)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateOrder(id int, data *models.Orders, db *gorm.DB) (models.Orders, error) {
	var order models.Orders
	err := db.Preload("Items").First(&order, id).Error
	if err != nil {
		return models.Orders{}, err
	}
	return order, err
}

func DeleteOrder(id int, db *gorm.DB) error {
	var order models.Orders

	del := db.Delete(&order, id)

	if del.Error != nil {
		return del.Error
	} else {
		return nil
	}

}
