package middlewares

import (
	"gorm.io/gorm"
	"technodom_test/db"
	"technodom_test/models"
)

func FindLinksByCategory(category string) ([]models.Link, error) {
	var links []models.Link
	result := db.DB.Model(&models.Link{}).Where("active_link LIKE ?", category+"%").Find(&links)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return links, nil
}
