package model

type Post struct {
	ID     int
	UserID int    `sql:"index"`
	Title  string `sql:"size:255"`
	Body   string `sql:"type:text"`
}

type Posts []Post
