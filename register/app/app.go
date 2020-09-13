package app

import (
	"os"

	"github.com/tabakazu/golang-webapi-demo/db"
	"github.com/tabakazu/golang-webapi-demo/server"
)

func Run() {
	os.Exit(func() int {
		d := db.NewPostgreSQLDB()
		s := server.NewServer(8080, d)

		termCh := make(chan os.Signal, 1)
		errCh := make(chan error, 1)
		go func() {
			errCh <- s.Start()
		}()

		select {
		case <-termCh:
			return 0
		case <-errCh:
			return 1
		}
	}())
}
