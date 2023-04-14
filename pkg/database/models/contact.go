package models

import (
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	Correspondent string `gorm:"primaryKey;unique;size:100" json:"correspondent"`
	// User          User   `gorm:"not null;" json:"user"`
	Address   string `gorm:"unique;not null;size:42" json:"address"`
	PublicKey []byte `gorm:"unique;not null" json:"publickey"`
	Channel   string `gorm:"unique;size:42" json:"channel"`
}
