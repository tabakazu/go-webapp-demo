package datastore

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type dbConfig struct {
	dsn string // データソースネーム
}

func NewDBConfig() *dbConfig {
	return &dbConfig{
		dsn: os.Getenv("MYSQL_URL"),
	}
}

func NewConnection(cfg *dbConfig) (*gorm.DB, func()) {
	if cfg == nil {
		cfg = &dbConfig{}
	}

	db, err := gorm.Open(mysql.Open(cfg.dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db, func() {
		if db != nil {
			sqlDB, err := db.DB()
			if err != nil {
				log.Fatal(err)
			}
			if err := sqlDB.Close(); err != nil {
				log.Fatal(err)
			}
		}
	}
}
