package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username string `json:"username" form:"username" xml:"username" gorm:"unique"`
	Password string `json:"password" form:"password" xml:"password"`
	Profile  string `json:"profile" form:"profile" xml:"profile"`
	Name     string `json:"name" form:"name" xml:"name"`
	Email    string `json:"email" form:"email" xml:"email"`
	Gender   string `json:"gender" form:"gender" xml:"gender"`
	Address  string `json:"address" form:"address" xml:"address"`
	Phone    string `json:"phone" form:"phone" xml:"phone"`
	Born     string `json:"born" form:"born" xml:"born"`
	Is_Admin bool   `gorm:"default:false" json:"is_admin" form:"is_admin" xml:"is_admin"`
}

type UserGetUsername struct {
	Username string `json:"username" form:"username" xml:"username"`
	Is_Admin bool   `json:"is_admin" form:"is_admin" xml:"is_admin"`
}

type RequestUsersLogin struct {
	Username string `json:"username" form:"username" xml:"username"`
	Password string `form:"password" json:"password"`
}

type RequestUsersRegister struct {
	Username         string `json:"username" form:"username" xml:"username" gorm:"unique"`
	Password         string `form:"password" json:"password"`
	Confirm_password string `form:"confirm_password" json:"confirm_password"`
	Is_Admin         bool   `gorm:"default:false" json:"is_admin" form:"is_admin" xml:"is_admin"`
}
