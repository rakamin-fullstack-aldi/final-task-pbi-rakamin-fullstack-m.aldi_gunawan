package models

import (
	"time"
)

type User struct {
	ID        uint    `gorm:"primaryKey"`
	Username  string  `gorm:"not null"`
	Email     string  `gorm:"unique;not null"`
	Password  string  `gorm:"not null"`
	Photos    []Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Authenticate(email, password string) (*User, error) {
	var user User
	err := DB.Where("email = ? AND password = ?", email, password).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
