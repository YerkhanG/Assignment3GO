package main

import (
	"Assignment3/Init"
	"Assignment3/models"
)

func init() {
	Init.ConnectDB()
	Init.LoadEnvVar()
}
func main() {
	Init.DB.AutoMigrate(&models.Book{})
}
