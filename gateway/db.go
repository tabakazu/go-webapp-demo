package gateway

import (
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

func NewDB(cfg *dbConfig) *gorm.DB {
	if cfg == nil {
		cfg = &dbConfig{}
	}

	db, err := gorm.Open(mysql.Open(cfg.dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
