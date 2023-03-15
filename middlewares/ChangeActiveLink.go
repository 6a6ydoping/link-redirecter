package middlewares

import (
	"fmt"
	"gorm.io/gorm"
	"technodom_test/models"
)

func ChangeActiveLink(db *gorm.DB, id uint, newActiveLink string) error {
	link := &models.Link{}
	if err := db.First(link, id).Error; err != nil {
		return err
	}
	fmt.Println(4)
	link.HistoryLink = link.ActiveLink
	link.ActiveLink = newActiveLink
	if err := db.Save(link).Error; err != nil {
		return err
	}
	fmt.Println(5)
	return nil
}
