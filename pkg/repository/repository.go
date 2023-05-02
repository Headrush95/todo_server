package repository

import (
	First_server "github.com/headrush95/first_server"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user First_server.User) (int, error)
	GetUser(username, password string) (First_server.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
