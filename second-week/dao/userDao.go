package userDao

import (
	"database/sql"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func FindUserById(id string) (*User, error) {
	// sql.FindOne(&user)
	return nil, sql.ErrNoRows
}
