package job

import (
	"leyra/app"
	"leyra/app/models"
)

type AuthenticateUser struct {
	Email    string
	Password string
}

func (au AuthenticateUser) Handle() model.User {
	// TODO: Validate these post fields here
	user := model.User{}

	app.S.DB.Where(
		"email = ? and password >= ?",
		au.Email,
		au.Password,
	).First(&user)

	return user
}
