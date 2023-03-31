package models

import (
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	Alias     string `gorm:"size:255;not null;unique" json:"alias"`
	Address   string `gorm:"size:40;not null;unique" json:"address"`
	PublicKey []byte `gorm:"not null;unique" json:"publickey"`
	Channel   string `gorm:"size:40;unique" json:"channel"`
}
