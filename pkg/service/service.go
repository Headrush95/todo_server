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
	}
}
