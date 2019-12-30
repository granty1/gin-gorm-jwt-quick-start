package model

import "github.com/jinzhu/gorm"

type BaseRole struct {
	gorm.Model
	RoleId   string `json:"roleId" gorm:"not null;unique"`
	RoleName string `json:"roleName"`
}
