package middlewares

import (
	"gorm.io/gorm"
	"technodom_test/models"
)

func GetAllLinks(db *gorm.DB, page, pageSize int) ([]models.Link, error) {
	var links []models.Link
	offset := (page - 1) * pageSize
	err := db.Table("links").Select("id, active_link, history_link").Offset(offset).Limit(pageSize).Order("id asc").Find(&links).Error
	if err != nil {
		return nil, err
	}
	return links, nil
}
