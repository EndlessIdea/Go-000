package dao

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

var (
	db      *sql.DB
	UserDao = &userDao{}
)

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type userDao struct{}

func (b *userDao) TableName() string {
	return "users"
}

func (b *userDao) GetByID(ctx context.Context, id int) (*User, error) {
	_, err := db.Query("select * from users where id = ?", id)
	return &User{}, errors.Wrap(err, "failed to get user in db")
}
