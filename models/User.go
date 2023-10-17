package models

import (
	"fmt"
	"time"
)

type User struct {
	UpdatedAt time.Time `gorm:"autoUpdateTime;column:updated_at"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:updated_at"`
	ID        uint      `gorm:"primaryKey;column:user_id"`
	Username  string    `gorm:"column:username;unique"`
	Password  string    `gorm:"column:password"`
}

func (u User) TableName() string {
	return fmt.Sprintf("users")
}
