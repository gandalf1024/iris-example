package models

import (
	"time"
)

type User struct {
	Id           int `orm:"column(id);auto" description:"用户ID"`
	Name         string
	Age          int
	Account      string
	Password     string
	PasswordSalt string
	CreateDate   time.Time
	Ip           string
	IsAdmin      int8
}

func (t *User) TableName() string {
	return "sys_user"
}
