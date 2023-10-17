package service

import (
	"auth/config"
	"auth/repository"
)

//go:generate go run github.com/golang/mock/mockgen -source=service.go -destination=service_mock.go -package=service

type Services struct {
	User User
}

type Deps struct {
	Repos *repository.Repositories
	Cgf   *config.Config
}

func NewServices(deps Deps) *Services {
	userService := NewUserService(deps.Repos.User)
	return &Services{User: userService}
}
