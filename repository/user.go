package repository

import (
	"auth/model"
	"gorm.io/gorm"
)

type User interface {
	Create(user model.User) (model.User, error)
	GetByID(id int) (model.User, error)
	GetByUsername(username string) (model.User, error)
}

type UserDB struct {
	db *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{
		db: db,
	}
}
func (r *UserDB) Create(user model.User) (model.User, error) {
	db := r.db
	q := db.Model(&model.User{})
	err := q.Save(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UserDB) GetByID(id int) (model.User, error) {
	return model.User{}, nil
}

func (r *UserDB) GetByUsername(username string) (model.User, error) {
	var user model.User
	db := r.db
	q := db.Model(&model.User{})
	err := q.Where("username = ?", username).First(&user).Error
	if err != nil {
		return model.User{}, nil
	}
	return user, nil
}
