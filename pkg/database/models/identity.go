package models

import (
	"gorm.io/gorm"
)

// TODO: secure private keys, RSA & ECDSA
type Identity struct {
	gorm.Model
	Alias string `gorm:"primaryKey;unique;size:100" json:"alias"`
	// Username string `gorm:"size:100" json:"username"`
	ECDSA string `gorm:"not null" json:"ecdsa"`
	RSA   []byte `gorm:"not null" json:"rsa"`
}
