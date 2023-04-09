package Init

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := "host=mel.db.elephantsql.com user=gjxdwwrp password=Ace2DILFUuJ3U39kXkOeOnqEdmvFE7aG dbname=gjxdwwrp port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Fail. DataBase")
	}
}
