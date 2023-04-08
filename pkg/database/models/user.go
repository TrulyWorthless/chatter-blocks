package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"primaryKey;unique;size:100" json:"username"`
	Password string `gorm:"not null;size:100" json:"password"`
}
