package db

import (
	"context"
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tabakazu/golang-webapi-demo/model"
)

type mysqlDB struct {
	db *sql.DB
}

func NewMySQLDB() *mysqlDB {
	db, _ := sql.Open("mysql", os.Getenv("MYSQL_URL"))
	return &mysqlDB{db: db}
}

func (m *mysqlDB) GetAllItems(ctx context.Context) ([]*model.Item, error) {
	rows, err := m.db.Query("SELECT BIN_TO_UUID(id, true), title FROM items")
	if err != nil {
		return nil, err
	}

	var result []*model.Item
	for rows.Next() {
		var item model.Item
		err = rows.Scan(&item.ID, &item.Title)
		if err != nil {
			break
		}
		result = append(result, &item)
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *mysqlDB) CreateItem(ctx context.Context, t *model.Item) error {
	stmt, err := m.db.Prepare("INSERT INTO items(id, title) VALUES( UUID_TO_BIN(?, true), ? )")
	if err != nil {
		stmt.Close()
		return err
	}

	if _, err := stmt.Exec(t.ID, t.Title); err != nil {
		return err
	}

	return nil
}
