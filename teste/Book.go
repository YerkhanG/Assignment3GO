package main

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Id          string
	Name        string
	Description string
	Price       string
}
