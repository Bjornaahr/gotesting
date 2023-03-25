package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type Users struct {
	gorm.Model
	Name     string
	DoB      time.Time
	Watchman bool
	Email    string `gorm:"unique"`
	Password string
}

func (user *Users) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
