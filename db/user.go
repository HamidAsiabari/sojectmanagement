package db

import (
	"gorm.io/gorm"
)

func (u User) Add() (result *gorm.DB) {

	result = DatabaseConn.db.Create(&u)
	return result
}

func (u *User) Get(id int) (result *gorm.DB) {
	result = DatabaseConn.db.First(&u, id)
	return
}

func (u *User) GetByUsername(username string) (result *gorm.DB) {
	result = DatabaseConn.db.Where("username= ?", username).First(&u)
	return
}

func (u *User) Update(column string, val string) {

	DatabaseConn.db.Model(&u).Where("username = ?", u.Username).Update(column, val)

}
