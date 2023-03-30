package models

import (
	"errors"
	"html"
	"strings"

	"gorm.io/gorm"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

// TODO: secure private keys, RSA & ECDSA
type Identity struct {
	Uuid     uint32    `gorm:"primary_key;unique" json:"uuid"`
	Alias    string    `gorm:"size:255;not null;unique" json:"alias"`
	Password string    `gorm:"size:100;not null;" json:"password"`
	ECDSA    string    `gorm:"size:64;not null;" json:"ecdsa"`
	RSA      [256]byte `gorm:"size:256;not null;" json:"rsa"`
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

func (i *Identity) Init() {
	i.Uuid = uuid.New().ID()
	i.Alias = html.EscapeString(strings.TrimSpace(i.Alias))
	i.Password = html.EscapeString(strings.TrimSpace(i.Password))
	i.ECDSA = html.EscapeString(strings.TrimSpace(i.ECDSA))
	i.RSA = [256]byte{}
}

func (i *Identity) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if i.Uuid == 0 {
			return errors.New("required ID")
		}
		if i.Alias == "" {
			return errors.New("required Alias")
		}
		if i.Password == "" {
			return errors.New("required Password")
		}
		if i.ECDSA == "" {
			return errors.New("required ECDSA")
		}
		if i.RSA == [256]byte{} {
			return errors.New("required RSA")
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

func (i *Identity) SaveIdentity(db *gorm.DB) (*Identity, error) {
	err := db.Debug().Create(&i).Error
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (i *Identity) FindAllIdentities(db *gorm.DB) (*[]Identity, error) {
	identities := []Identity{}
	err := db.Debug().Model(&Identity{}).Limit(100).Find(&identities).Error
	if err != nil {
		return nil, err
	}

	return &identities, nil
}

func (i *Identity) FindIdentityByID(db *gorm.DB, uuid uint32) (*Identity, error) {
	err := db.Debug().Model(Identity{}).Where("uuid = ?", uuid).Take(&i).Error
	if err != nil {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("identity not found")
	}

	return i, nil
}

func (i *Identity) FindIdentityByAlias(db *gorm.DB, alias string) (*Identity, error) {
	err := db.Debug().Model(Identity{}).Where("alias = ?", alias).Take(&i).Error
	if err != nil {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("identity not found")
	}

	return i, nil
}

func (i *Identity) UpdateAnIdentity(db *gorm.DB, uuid uint32) (*Identity, error) {
	err := i.BeforeSave()
	if err != nil {
		panic(err)
	}

	db = db.Debug().Model(&Identity{}).Where("uuid = ?", uuid).Take(&Identity{}).UpdateColumns(
		map[string]interface{}{
			"alias":    i.Alias,
			"password": i.Password,
			"ecdsa":    i.ECDSA,
			"rsa":      i.RSA,
		},
	)

	if db.Error != nil {
		return nil, db.Error
	}

	err = db.Debug().Model(&Identity{}).Where("uuid = ?", uuid).Take(&i).Error
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (i *Identity) DeleteAnIdentity(db *gorm.DB, uuid uint32) (int64, error) {
	db = db.Debug().Model(&Identity{}).Where("uuid = ?", uuid).Take(&Identity{}).Delete(&Identity{})
	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
