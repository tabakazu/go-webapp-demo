package webapi

import (
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type context struct {
	echo.Context
}

func (c context) ApiSessionUserID() int {
	claims := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)
	userID, err := strconv.Atoi(claims["user_id"].(string))
	if err != nil {
		return 0
	}
	return userID
}
