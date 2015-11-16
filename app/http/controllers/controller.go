package controller

import (
	"gopkg.in/leyra/echo.v1"

	"leyra/app"
	"leyra/app/models"
)

type Controller struct {
}

func CurrentUser(c *echo.Context) model.User {
	user := model.User{}
	app.S.DB.Where("id = ?", app.S.Get(c, "user")).First(&user)

	return user
}

func IsAuthenticated(c *echo.Context) bool {
	if CurrentUser(c).ID > 0 {
		return true
	}

	return false
}
