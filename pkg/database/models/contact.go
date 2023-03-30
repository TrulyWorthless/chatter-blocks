package models

type Contact struct {
	Uuid      uint32   `gorm:"primary_key;" json:"uuid"`
	Alias     string   `gorm:"size:255;not null;" json:"alias"`
	Address   string   `gorm:"size:40;not null;unique" json:"address"`
	PublicKey [64]byte `gorm:"size:64;not null;unique" json:"publickey"`
	Channel   string   `gorm:"size:40;unique" json:"channel"`
}
