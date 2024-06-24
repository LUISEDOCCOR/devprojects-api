package models

import "gorm.io/gorm"

type Projects struct {
	gorm.Model

	Title       string `gorm:"not null" json:"title"`
	Content     string `gorm:"not null" json:"content"`
	Id_category uint   `gorm:"not null" json:"id_category"`
}
