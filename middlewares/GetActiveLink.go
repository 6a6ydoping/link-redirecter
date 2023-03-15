package middlewares

import (
	"errors"
	"gorm.io/gorm"
	"technodom_test/models"
)

func FindActiveLink(DB *gorm.DB, historyLink string) (string, error) {
	var link models.Link
	err := DB.Where("history_link = ?", historyLink).First(&link).Error
	if err != nil {
		return "", err
	}

	if link.ActiveLink != "" {
		return link.ActiveLink, nil
	} else {
		return "", errors.New("no active link found")
	}
}
