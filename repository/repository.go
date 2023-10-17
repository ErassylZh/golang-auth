package repository

import (
	"auth/config"
	"gorm.io/gorm"
)

type Repositories struct {
	User User
}

func NewRepositories(db *gorm.DB, cfg *config.Config) (*Repositories, error) {
	userRepo := NewUserDB(db)
	return &Repositories{User: userRepo}, nil
}
