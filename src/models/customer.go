package models

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID       uint    `json:"id,omitempty"`
	Fullname string  `json:"fullname" binding:"required"`
	Email    string  `json:"email" binding:"required,email"`
	Gender   string  `json:"gender" binding:"required"`
	DOB      *string `json:"dob"`
	Phone    string  `json:"phone"`
	Address  string  `json:"address"`
}
