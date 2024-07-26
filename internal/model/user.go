package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:default:null json:"first_name,omitempty"`
	LastName  string `gorm:default:null json:"last_name,omitempty"`
	Email     string `gorm:"unique" json:"email"`
}

type UserCreateForm struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string
	Email     string
}
