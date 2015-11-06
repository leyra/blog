package job

import (
	"time"

	"leyra/app"
	"leyra/app/models"
)

type CreateUser struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (cu CreateUser) Handle() model.User {
	// TODO: Validate these post fields here
	user := model.User{
		FirstName: cu.FirstName,
		LastName:  cu.LastName,
		Email:     cu.Email,
		Password:  cu.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db := app.S.DB
	db.NewRecord(user)
	db.Create(&user)

	return user
}
