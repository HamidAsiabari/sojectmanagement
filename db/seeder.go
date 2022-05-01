package db

import (
	"github.com/HamidAsiabari/sojectmanagement/validation"
)

func addBaseData() {
	addAdmin()
}

func addAdmin() {

	u := User{
		Username:  "Admin",
		Password:  "Admin22AA@#!",
		FirstName: "Admin",
		LastName:  "",
		Email:     "admin@admin.com",
	}
	var err error
	u.Password, err = validation.HashPassword(u.Password)
	if err != nil {
		println("Can not create Admin user!")
		return
	}
	res := u.Add()
	if res.Error != nil {
		println("Can not create Admin user!")
	}
}
