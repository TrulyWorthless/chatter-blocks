package models

import (
	"errors"
	"html"
	"strings"

	"gorm.io/gorm"
)

type Contact struct {
	Uuid      uint32   `gorm:"primary_key;" json:"uuid"`
	Alias     string   `gorm:"size:255;not null;" json:"alias"`
	Address   string   `gorm:"size:40;not null;unique" json:"address"`
	PublicKey [64]byte `gorm:"size:64;not null;unique" json:"publickey"`
	Channel   string   `gorm:"size:40;unique" json:"channel"`
}

func (c *Contact) Init() {
	c.Uuid = 0
	c.Alias = html.EscapeString(strings.TrimSpace(c.Alias))
	c.Address = html.EscapeString(strings.TrimSpace(c.Address))
	c.Channel = html.EscapeString(strings.TrimSpace(c.Channel))
	c.PublicKey = [64]byte{}
}

func (c *Contact) Validate() error {
	if c.Uuid == 0 {
		return errors.New("required ID")
	}
	if c.Alias == "" {
		return errors.New("required Alias")
	}
	if c.Address == "" {
		return errors.New("required Address")
	}
	if c.Channel == "" {
		return errors.New("required Channel")
	}
	if c.PublicKey == [64]byte{} {
		return errors.New("required PublicKey")
	}
	return nil
}

func (c *Contact) SaveContact(db *gorm.DB) (*Contact, error) {
	err := db.Debug().Model(&Contact{}).Create(&c).Error
	if err != nil {
		return nil, err
	}

	if c.Uuid != 0 {
		err = db.Debug().Model(&Contact{}).Where("uuid = ?", c.Uuid).Take(&c).Error
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Contact) FindAllContacts(db *gorm.DB) (*[]Contact, error) {
	contacts := []Contact{}
	err := db.Debug().Model(&Contact{}).Limit(100).Find(&contacts).Error
	if err != nil {
		return nil, err
	}

	return &contacts, nil
}

func (c *Contact) FindContactByID(db *gorm.DB, uuid uint32) (*Contact, error) {
	err := db.Debug().Model(Contact{}).Where("uuid = ?", uuid).Take(&c).Error
	if err != nil {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("contact not found")
	}

	return c, nil
}

func (c *Contact) UpdateAContact(db *gorm.DB, uuid uint32) (*Contact, error) {
	db = db.Debug().Model(&Contact{}).Where("uuid = ?", uuid).Take(&Contact{}).UpdateColumns(
		map[string]interface{}{
			"alias":     c.Alias,
			"address":   c.Address,
			"publickey": c.PublicKey,
			"channel":   c.Channel,
		},
	)

	if db.Error != nil {
		return nil, db.Error
	}

	err := db.Debug().Model(&Contact{}).Where("uuid = ?", uuid).Take(&c).Error
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Contact) DeleteAContact(db *gorm.DB, uuid uint32) (int64, error) {
	db = db.Debug().Model(&Contact{}).Where("uuid = ?", uuid).Take(&Contact{}).Delete(&Contact{})
	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
