package model

import (
	"time"
)

type User struct {
	ID        int
	Email     string `sql:"size:255"`
	Password  string `sql:"size:255"`
	FirstName string `sql:"size:255"`
	LastName  string `sql:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Posts []Post
}
