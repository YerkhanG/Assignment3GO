package main

import "gorm.io/gorm"

type Admin struct {
	Id    string
	Name  string
	login string
}

func (A Admin) getBookById(id int, db *gorm.DB) {
	db.First(&)
}
