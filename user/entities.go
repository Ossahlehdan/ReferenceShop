package user

import (
	"github.com/jinzhu/gorm"
)

//User represent an user
type User struct {
	gorm.Model
	FirstName string `gorm:"size:255;not null" json:"firstName"`
	LastName  string `gorm:"size:255;not null" json:"lastName"`
	Email     string `gorm:"size:100;not null;unique" json:"email"`
	Password  string `gorm:"size:100;not null;" json:"password"`
}

//LoginInfo represent the login info
type LoginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
