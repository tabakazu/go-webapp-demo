package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type connection struct {
	dbURL  string
	config *gorm.Config
}

func New(dbURL string) *connection {
	return &connection{dbURL: dbURL, config: &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}}
}

func (c *connection) SilentMode() {
	c.config.Logger = logger.Default.LogMode(logger.Silent)
}

func (c *connection) Connect() *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.dbURL), c.config)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
