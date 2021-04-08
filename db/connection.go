package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnection(dbURL string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
