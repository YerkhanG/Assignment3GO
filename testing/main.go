package testing

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.AutoMigrate(&Book{})

}
