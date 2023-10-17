package repository

import (
	"auth/config"
	"gorm.io/gorm"
)

type Repositories struct {
}

func NewRepositories(db *gorm.DB, cfg *config.Config) (*Repositories, error) {
	return &Repositories{}, nil
}
