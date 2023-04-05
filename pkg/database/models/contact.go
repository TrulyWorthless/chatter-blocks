package models

import (
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	Alias     string `gorm:"size:255;not null;unique" json:"alias"`
	Address   string `gorm:"size:42;not null;unique" json:"address"`
	PublicKey []byte `gorm:"not null;unique" json:"publickey"`
	Channel   string `gorm:"size:42;unique" json:"channel"`
}
