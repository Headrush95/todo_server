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
	Create(userId int, list First_server.TodoList) (int, error)
	GetAll(userId int) ([]First_server.TodoList, error)
	GetById(userId, listId int) (First_server.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input First_server.UpdateListInput) error
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
		TodoList:      NewTodoListPostgres(db),
	}
}
