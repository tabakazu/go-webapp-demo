package gateway_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tabakazu/golang-webapi-demo/db"
	"github.com/tabakazu/golang-webapi-demo/domain"
	"github.com/tabakazu/golang-webapi-demo/gateway"
)

func TestItemRepository_Find(t *testing.T) {
	t.Parallel()

	d := db.NewConnection(os.Getenv("MYSQL_TEST_URL"))

	// START TRANSACTION
	tx := d.Begin()

	// INSERT TEST DATA
	tx.Create(&domain.Item{ID: domain.ItemID(1), Name: "aaa", Amount: 100})
	tx.Create(&domain.Item{ID: domain.ItemID(2), Name: "bbb", Amount: 200})

	// TEST
	itemRepo := gateway.NewItemRepository(tx)
	item, err := itemRepo.Find(domain.ItemID(1))
	assert.NoError(t, err)
	assert.Equal(t, item.Name, "aaa")

	item, err = itemRepo.Find(domain.ItemID(2))
	assert.NoError(t, err)
	assert.Equal(t, item.Name, "bbb")

	_, err = itemRepo.Find(domain.ItemID(3))
	assert.Error(t, err)

	// ROLLBACK TEST DATA
	tx.Rollback()
}
