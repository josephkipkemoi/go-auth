package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	PhoneNumber int64 `gorm:"unique;not null;" json:"phoneNumber"`
	Password string `gorm:"size:255;not null;" json:"password"`
}	

func (u *User) SaveUser() (*User, error) {
	var err error
	err = DB.Create(&u).Error

	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) HashPassword() error {
	hp, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hp)

	return nil
}

func (u *User) AuthUser() *gorm.DB {
	tx := DB.Where("phone_number",u.PhoneNumber).First(u)
	return tx
}