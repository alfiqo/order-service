package models

import (
	"order-service/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `json:"id,omitempty"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(user *User, password string) (token string, err error) {

	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err = utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil

}
