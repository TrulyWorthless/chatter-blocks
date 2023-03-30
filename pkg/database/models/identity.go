package models

// TODO: secure private keys, RSA & ECDSA
type Identity struct {
	Uuid     uint32    `gorm:"primary_key;unique" json:"uuid"`
	Alias    string    `gorm:"size:255;not null;unique" json:"alias"`
	Password string    `gorm:"size:100;not null;" json:"password"`
	ECDSA    string    `gorm:"size:64;not null;" json:"ecdsa"`
	RSA      [256]byte `gorm:"size:256;not null;" json:"rsa"`
}
