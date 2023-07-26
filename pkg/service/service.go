package service

import (
	first_server "github.com/headrush95/first_server"
	"github.com/headrush95/first_server/pkg/repository"
)

type Authorization interface {
	CreateUser(user first_server.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list first_server.TodoList) (int, error)
	GetAll(userId int) ([]first_server.TodoList, error)
	GetById(userId, listId int) (first_server.TodoList, error)
	Delete(userId, ListId int) error
	Update(userId, listId int, input first_server.UpdateListInput) error
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
