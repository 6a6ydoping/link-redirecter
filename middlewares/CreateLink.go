package middlewares

import (
	"gorm.io/gorm"
	"technodom_test/models"
)

func CreateLink(DB *gorm.DB, activeLink, historyLink string) error {
	link := &models.Link{
		ActiveLink:  activeLink,
		HistoryLink: historyLink,
	}
	result := DB.Create(link)
	return result.Error
}
