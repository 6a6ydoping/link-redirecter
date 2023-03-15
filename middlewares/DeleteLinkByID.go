package middlewares

import (
	"gorm.io/gorm"
	"technodom_test/models"
)

func DeleteLinkById(DB *gorm.DB, id uint) error {
	err := DB.Delete(&models.Link{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
