package middlewares

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"technodom_test/models"
)

func GetLinkByID(db *gorm.DB, id int) (*models.Link, error) {
	link := &models.Link{}
	result := db.First(link, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Link with ID %d not found", id)
		}
		return nil, result.Error
	}
	return link, nil
}
