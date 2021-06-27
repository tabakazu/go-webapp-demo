package webapi

import (
	"net/http"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tabakazu/go-webapp/interfaces/webapi/controller"
)

func CustomContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &customContext{c, nil}
		return next(cc)
	}
}

func UserTokenAuth(next echo.HandlerFunc) echo.HandlerFunc {
	authMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
		ErrorHandler: func(err error) error {
			return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
		},
	})

	return authMiddleware(func(c echo.Context) error {
		cc := c.(*customContext)

		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			if err := next(cc); err != nil {
				c.Error(err)
			}
			return nil
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			if err := next(cc); err != nil {
				c.Error(err)
			}
			return nil
		}

		userIDstring, ok := claims["user_id"].(string)
		if !ok {
			if err := next(cc); err != nil {
				c.Error(err)
			}
			return nil
		}

		userID, err := strconv.Atoi(userIDstring)
		if err != nil {
			if err := next(cc); err != nil {
				c.Error(err)
			}
			return nil
		}

		cc.session = &controller.Session{UserID: userID}
		if err := next(cc); err != nil {
			c.Error(err)
		}
		return nil
	})
}
