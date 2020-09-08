package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/tabakazu/golang-webapi-demo/model"
)

type postgresqlDB struct {
	db *sql.DB
}

func NewPostgreSQLDB() *postgresqlDB {
	db, _ := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	return &postgresqlDB{db: db}
}

func (m *postgresqlDB) GetAllItems(ctx context.Context) ([]*model.Item, error) {
	rows, err := m.db.Query("SELECT id, title FROM items")
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

func (m *postgresqlDB) PutItem(ctx context.Context, t *model.Item) error {
	stmt, err := m.db.Prepare("INSERT INTO items (id, title) VALUES ( $1, $2 )")
	if err != nil {
		fmt.Println(err)
		stmt.Close()
		return err
	}

	if _, err := stmt.Exec(t.ID, t.Title); err != nil {
		return err
	}

	return nil
}
