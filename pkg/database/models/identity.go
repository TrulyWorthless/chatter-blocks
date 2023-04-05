package models

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// TODO: secure private keys, RSA & ECDSA
type Identity struct {
	gorm.Model
	Alias    string `gorm:"size:255;not null;unique" json:"alias"`
	Password string `gorm:"size:100;not null;" json:"password"`
	ECDSA    string `gorm:"not null;" json:"ecdsa"`
	RSA      []byte `gorm:"not null;" json:"rsa"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func (i *Identity) BeforeSave() error {
	hashedPassword, err := Hash(i.Password)
	if err != nil {
		return err
	}

	i.Password = string(hashedPassword)
	return nil
}

func (i *Identity) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if i.Alias == "" {
			return errors.New("required Alias")
		}
		if i.Password == "" {
			return errors.New("required Password")
		}
		if i.ECDSA == "" {
			return errors.New("required ECDSA")
		}
		return nil
	case "login":
		if i.Alias == "" {
			return errors.New("required Alias")
		}
		if i.Password == "" {
			return errors.New("required Password")
		}
		return nil
	default:
		return errors.New("invalid Use Case")
	}
}
