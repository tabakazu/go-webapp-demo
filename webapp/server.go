package webapp

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tabakazu/golang-webapi-demo/webapp/controller"
)

type RoutingSet struct {
	Items controller.Items
}

type server struct {
	*echo.Echo
}

func NewServer(r RoutingSet) *server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	e.GET("/items", r.Items.List)
	e.POST("/items", r.Items.Create)
	e.GET("/items/:id", r.Items.Show)
	e.PUT("/items/:id", r.Items.Update)
	e.DELETE("/items/:id", r.Items.Delete)

	return &server{e}
}

func (s server) ListenAndServe() {
	s.Logger.Fatal(s.Start(":8080"))
}
