package server

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/tabakazu/golang-webapi-demo/db"
)

type Server struct {
	server *http.Server
}

func NewServer(port int, d db.DB) *Server {
	mux := http.NewServeMux()
	mux.Handle("/items", &itemsHandler{db: d})

	return &Server{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: mux,
		},
	}
}

func (s *Server) Start() error {
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}

func Run() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	termCh := make(chan os.Signal, 1)
	// d := db.NewMemoryDB()
	d := db.NewMySQLDB()
	s := NewServer(8080, d)
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
}
