package models

import "gorm.io/gorm"

type Categories struct {
	gorm.Model

	Name string `gorm:"not null" json:"name"`
}
