package model

import (
	"time"
)

type User struct{
	Id int `gorm:"primary_key"`
	Name string `gorm:"type:varchar(30);"`
	Pass string `gorm:"type:varchar(64);"`
	Addtime time.Time 
}

func UserList() []User {

	var userlist []User

	db.Find(&userlist)

	return userlist
}
