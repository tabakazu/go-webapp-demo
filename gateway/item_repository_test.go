package gateway_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tabakazu/golang-webapi-demo/db"
	"github.com/tabakazu/golang-webapi-demo/domain"
	"github.com/tabakazu/golang-webapi-demo/gateway"
)

func TestItemRepository_FindAll(t *testing.T) {
	t.Parallel()

	// CONNECT DB
	d := db.NewConnection(os.Getenv("MYSQL_TEST_URL"))

	// START TRANSACTION
	tx := d.Begin()

	// INSERT TEST DATA
	tx.Create(&domain.Item{ID: domain.ItemID(1), Name: "aaa", Amount: 100})
	tx.Create(&domain.Item{ID: domain.ItemID(2), Name: "bbb", Amount: 200})

	// TEST
	itemRepo := gateway.NewItemRepository(tx)
	items, err := itemRepo.FindAll()
	assert.NoError(t, err)
	assert.Equal(t, len(items), 2)

	// ROLLBACK TEST DATA
	tx.Rollback()
}

func TestItemRepository_Find(t *testing.T) {
	t.Parallel()

	// CONNECT DB
	d := db.NewConnection(os.Getenv("MYSQL_TEST_URL"))

	// START TRANSACTION
	tx := d.Begin()

	// INSERT TEST DATA
	tx.Create(&domain.Item{ID: domain.ItemID(1), Name: "aaa", Amount: 100})

	// TEST
	itemRepo := gateway.NewItemRepository(tx)
	item, err := itemRepo.Find(domain.ItemID(1))
	assert.NoError(t, err)
	assert.Equal(t, item.ID, domain.ItemID(1))

	// ROLLBACK TEST DATA
	tx.Rollback()
}

func TestItemRepository_Create(t *testing.T) {
	t.Parallel()

	// CONNECT DB
	d := db.NewConnection(os.Getenv("MYSQL_TEST_URL"))

	// START TRANSACTION
	tx := d.Begin()

	// TEST
	itemRepo := gateway.NewItemRepository(tx)
	var item = domain.Item{Name: "NewItem", Amount: 1000}
	err := itemRepo.Create(&item)
	assert.NoError(t, err)

	// ROLLBACK TEST DATA
	tx.Rollback()
}

func TestItemRepository_UpdateAttributes(t *testing.T) {
	t.Parallel()

	// CONNECT DB
	d := db.NewConnection(os.Getenv("MYSQL_TEST_URL"))

	// START TRANSACTION
	tx := d.Begin()

	// INSERT TEST DATA
	tx.Create(&domain.Item{ID: domain.ItemID(1), Name: "aaa", Amount: 100})

	// TEST
	itemRepo := gateway.NewItemRepository(tx)
	item, _ := itemRepo.Find(domain.ItemID(1))
	attributes := map[string]interface{}{
		"name":   "NewName",
		"amount": 0,
	}
	err := itemRepo.UpdateAttributes(&item, attributes)
	assert.NoError(t, err)
	assert.Equal(t, item.Name, attributes["name"])
	assert.Equal(t, item.Amount, attributes["amount"])

	// ROLLBACK TEST DATA
	tx.Rollback()
}

func TestItemRepository_SoftDelete(t *testing.T) {
	t.Parallel()

	// CONNECT DB
	d := db.NewConnection(os.Getenv("MYSQL_TEST_URL"))

	// START TRANSACTION
	tx := d.Begin()

	// INSERT TEST DATA
	tx.Create(&domain.Item{ID: domain.ItemID(1), Name: "aaa", Amount: 100})

	// TEST
	itemRepo := gateway.NewItemRepository(tx)
	item, _ := itemRepo.Find(domain.ItemID(1))
	err := itemRepo.SoftDelete(&item)
	assert.NoError(t, err)

	// ROLLBACK TEST DATA
	tx.Rollback()
}
