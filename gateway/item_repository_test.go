package gateway_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tabakazu/golang-webapi-demo/db"
	"github.com/tabakazu/golang-webapi-demo/domain"
	"github.com/tabakazu/golang-webapi-demo/gateway"
)

func newItemRepository() gateway.ItemRepository {
	conn := db.NewConnection(os.Getenv("MYSQL_TEST_URL"))
	tx := conn.Begin()
	return gateway.NewItemRepository(tx)
}

func TestFindAllItems(t *testing.T) {
	r := newItemRepository()
	defer r.DB.Rollback()

	t.Run("is expected return all items", func(t *testing.T) {
		r.Create(&domain.Item{})
		r.Create(&domain.Item{})
		items, _ := r.FindAll()
		assert.Equal(t, len(items), 2)
	})
}

func TestFindItem(t *testing.T) {
	r := newItemRepository()
	defer r.DB.Rollback()

	t.Run("is expected return item with specified id", func(t *testing.T) {
		newItem := domain.Item{}
		r.Create(&newItem)
		item, _ := r.Find(newItem.ID)
		assert.Equal(t, item, newItem)
	})
	t.Run("is expected return error with record not found", func(t *testing.T) {
		id := domain.ItemID(1)
		_, err := r.Find(id)
		assert.NotNil(t, err)
	})
}

func TestCreateItem(t *testing.T) {
	r := newItemRepository()
	defer r.DB.Rollback()

	t.Run("is expected increase record count by 1", func(t *testing.T) {
		newItem := domain.Item{}
		err := r.Create(&newItem)
		assert.NoError(t, err)
		item, _ := r.Find(newItem.ID)
		assert.NotNil(t, item)
	})
}

func TestUpdateItemByAttributes(t *testing.T) {
	r := newItemRepository()
	defer r.DB.Rollback()

	t.Run("is expected change oldName to newName", func(t *testing.T) {
		item := domain.Item{Name: "oldName"}
		r.Create(&item)
		item, _ = r.Find(item.ID)
		assert.Equal(t, item.Name, "oldName")

		err := r.UpdateAttributes(&item, map[string]interface{}{"name": "oldName"})
		item, _ = r.Find(item.ID)
		assert.Equal(t, item.Name, "oldName")
		assert.NoError(t, err)
	})
}

func TestDeleteItem(t *testing.T) {
	r := newItemRepository()
	defer r.DB.Rollback()

	t.Run("is expected Delete() to change count by -1", func(t *testing.T) {
		item := domain.Item{}
		r.Create(&item)
		items, _ := r.FindAll()
		assert.Equal(t, len(items), 1)

		r.Delete(&item)
		items, _ = r.FindAll()
		assert.Equal(t, len(items), 0)
	})
}
