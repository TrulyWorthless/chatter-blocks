package seed

import (
	"github.com/trulyworthless/chatter-blocks/pkg/api/models"
	"gorm.io/gorm"
)

var identities = []models.Identity{
	{
		Uuid:     0,
		Alias:    "Alice",
		Password: "password0",
		ECDSA:    "4dcaef613ec559c7c86e6869297a2f4db360074fccfda092c47118d023692d29",
		RSA:      [256]byte{100},
	},
	{
		Uuid:     1,
		Alias:    "Bob",
		Password: "password1",
		ECDSA:    "59dea86fb1a061e1c18b63dbe0659101093a270a28a697144d82fc85841b07a9",
		RSA:      [256]byte{101},
	},
	{
		Uuid:     2,
		Alias:    "Victor",
		Password: "password2",
		ECDSA:    "80f9654b3b7ea77d824118677d6380b9fd70004b9d4a1f06d3c03ec29d758b36",
		RSA:      [256]byte{102},
	},
	{
		Uuid:     3,
		Alias:    "Sybil",
		Password: "password3",
		ECDSA:    "864a47b5b78b71cb1ae4afaf9d0f4679a1317646b4599d600d751bc3433f4dc8",
		RSA:      [256]byte{103},
	},
}

var contacts = []models.Contact{
	{
		Uuid:      00,
		Alias:     "Alice",
		Address:   "0xE5b4C8bcA8237D9F2D7201f7bC744b697aCF8A23",
		PublicKey: [64]byte{200},
		Channel:   "0x46292f6bA764E34f69774913EEA2F626591cb1E9",
	},
	{
		Uuid:      11,
		Alias:     "Bob",
		Address:   "0xe8148308ef1692e0f169F3FB8C0608a6E9625603",
		PublicKey: [64]byte{201},
		Channel:   "0x46292f6bA764E34f69774913EEA2F626591cb1E9",
	},
	{
		Uuid:      22,
		Alias:     "Victor",
		Address:   "0x89B946e790f5fAda7dB80950B9327A8E3BEA330f",
		PublicKey: [64]byte{202},
		Channel:   "0x46292f6bA764E34f69774913EEA2F626591cb1E9",
	},
	{
		Uuid:      33,
		Alias:     "Sybil",
		Address:   "0x9a548847E54225E354f07cdD2291a3A4a1bdF060",
		PublicKey: [64]byte{203},
		Channel:   "0x46292f6bA764E34f69774913EEA2F626591cb1E9",
	},
}

func Load(db *gorm.DB) {
	if db.Migrator().HasTable(&models.Identity{}) {
		err := db.Migrator().DropTable(&models.Identity{})
		if err != nil {
			panic(err)
		}
	}

	if db.Migrator().HasTable(&models.Contact{}) {
		err := db.Migrator().DropTable(&models.Contact{})
		if err != nil {
			panic(err)
		}
	}

	err := db.AutoMigrate(&models.Identity{}, &models.Contact{})
	if err != nil {
		panic((err))
	}

	for i := range identities {
		err = db.Debug().Model(&models.Identity{}).Create(&identities[i]).Error
		if err != nil {
			panic(err)
		}

		err = db.Debug().Model(&models.Contact{}).Create(&contacts[i]).Error
		if err != nil {
			panic(err)
		}
	}
}
