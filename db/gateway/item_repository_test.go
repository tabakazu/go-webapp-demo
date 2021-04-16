package gateway_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tabakazu/golang-webapi-demo/db"
	"github.com/tabakazu/golang-webapi-demo/domain/entity"
	"github.com/tabakazu/golang-webapi-demo/domain/value"
	"github.com/tabakazu/golang-webapi-demo/db/gateway"
)

func newItemRepository() gateway.ItemRepository {
	d := db.New(os.Getenv("MYSQL_TEST_URL"))
	d.SilentMode()
	conn := d.Connect()
	tx := conn.Begin()
	return gateway.NewItemRepository(tx)
}

func TestFindAllItems(t *testing.T) {
	r := newItemRepository()
	defer r.DB.Rollback()

	t.Run("is expected return all items", func(t *testing.T) {
		r.Create(&entity.Item{})
		r.Create(&entity.Item{})
		items, _ := r.FindAll()
		assert.Equal(t, len(items), 2)
	})
}

func TestFindItem(t *testing.T) {
	r := newItemRepository()
	defer r.DB.Rollback()

	t.Run("is expected return item with specified id", func(t *testing.T) {
		newItem := entity.Item{}
		r.Create(&newItem)
		item, _ := r.Find(newItem.ID)
		assert.Equal(t, item, newItem)
	})
	t.Run("is expected return error with record not found", func(t *testing.T) {
		id := value.ItemID(1)
		_, err := r.Find(id)
		assert.NotNil(t, err)
	})
}

func TestCreateItem(t *testing.T) {
	r := newItemRepository()
	defer r.DB.Rollback()

	t.Run("is expected increase record count by 1", func(t *testing.T) {
		newItem := entity.Item{}
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
		item := entity.Item{Name: "oldName"}
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
		item := entity.Item{}
		r.Create(&item)
		items, _ := r.FindAll()
		assert.Equal(t, len(items), 1)

		r.Delete(&item)
		items, _ = r.FindAll()
		assert.Equal(t, len(items), 0)
	})
}
