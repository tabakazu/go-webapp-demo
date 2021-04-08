package db

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type connection *gorm.DB

func NewConnection() connection {
	db, err := gorm.Open(mysql.Open(os.Getenv("MYSQL_URL")), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	return connection(db)
}
