package models

type Link struct {
	ID          uint `gorm:"primaryKey"`
	ActiveLink  string
	HistoryLink string
}
