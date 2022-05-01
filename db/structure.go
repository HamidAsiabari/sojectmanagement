package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username            string `gorm:"not null,unique"`
	Password            string `gorm:"not null"`
	FirstName           string
	LastName            string
	Email               string
	ChangePasswordToken string
}
type Role struct {
	gorm.Model
	Name     string
	Admin    bool
	Accesses []Access `gorm:"many2many:role_acceses;"`
}

type Access struct {
	gorm.Model
	Name string
	Code string
}

type Company struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	OwnerId     int
}
type Tag struct {
	gorm.Model
	Name string `gorm:"not null"`
}
