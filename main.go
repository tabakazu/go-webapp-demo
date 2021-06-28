package main

import (
	"github.com/tabakazu/go-webapp/external/datastore"
)

func main() {
	db, dbClose := datastore.NewConnection(datastore.NewDBConfig())
	defer dbClose()

	srv := InitializeServer(db)
	srv.ListenAndServe()
}
