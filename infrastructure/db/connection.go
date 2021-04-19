package db

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type connection struct {
	dbURL  string
	config *gorm.Config
}

func NewConnection() *gorm.DB {
	c := New(os.Getenv("MYSQL_URL"))
	return c.Connect()
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
