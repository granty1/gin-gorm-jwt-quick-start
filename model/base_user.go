package model

import (
	"github.com/gin-cli/tools"
	"github.com/jinzhu/gorm"
)

import uuid "github.com/satori/go.uuid"

//BaseUser system base user model
type BaseUser struct {
	gorm.Model
	UUID     uuid.UUID `json:"uuid"`
	Username string    `json:"userName" uri:"name"`
	Password string    `json:"-"`
	Email    string    `json:"email"`
	NickName string    `json:"nickName"`
	Phone    string    `json:"phone"`
	RoleID   string    `json:"roleId" uri:"roleId"`
	Role     BaseRole  `json:"role" gorm:"ForeignKey:RoleId;AssociationForeignKey:RoleId"`
}

func (u *BaseUser) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("UUID", uuid.NewV1())
}

func GetUserList() []BaseUser {
	var users []BaseUser
	connect.mysqlConnect.Preload("Role").Find(&users)
	return users
}

func GetUser(username, password string) (*BaseUser, bool) {
	var user BaseUser
	err := connect.mysqlConnect.Preload("Role").Where(&BaseUser{
		Username: username,
		Password: tools.Md5(password),
	}).First(&user).Error
	if err != nil {
		return nil, false
	}
	return &user, true
}

//func (u *BaseUser) Login() (BaseUser, bool) {
//	var user BaseUser
//	mysql.Instance().Where("")
//}
