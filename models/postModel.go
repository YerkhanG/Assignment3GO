package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID          uint
	Title       string
	Description string
	Cost        uint
}
