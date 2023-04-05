package handlers

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gofiber/fiber/v2"
	"github.com/trulyworthless/chatter-blocks/pkg/crypt"
	"github.com/trulyworthless/chatter-blocks/pkg/database"
	"github.com/trulyworthless/chatter-blocks/pkg/database/models"
	"github.com/trulyworthless/chatter-blocks/pkg/web3"
)

func ExportPublicKeyById(c *fiber.Ctx) error {
	id := c.Params("id")
	identity := models.Identity{}
	result := database.DB.Db.First(&identity, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	crypt.GenerateRSAPublicKeyFile(identity.Alias, &crypt.GetRSAPrivateKeyFromBytes(identity.RSA).PublicKey)

	return c.Status(200).JSON(identity)
}

func ExportPublicKeyByAlias(c *fiber.Ctx) error {
	alias := c.Params("alias")
	identity := models.Identity{}
	result := database.DB.Db.Where("alias = ?", alias).First(&identity)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	crypt.GenerateRSAPublicKeyFile(identity.Alias, &crypt.GetRSAPrivateKeyFromBytes(identity.RSA).PublicKey)

	return c.Status(200).JSON(identity)
}

// TODO use common.Address
func ExportAddressById(c *fiber.Ctx) error {
	id := c.Params("id")
	identity := models.Identity{}
	result := database.DB.Db.First(&identity, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	p, _ := crypto.HexToECDSA(identity.ECDSA)
	web3.GenerateBlockchainAddressFile(identity.Alias, crypto.PubkeyToAddress(p.PublicKey))

	return c.Status(200).JSON(identity)
}

func ExportAddressByAlias(c *fiber.Ctx) error {
	alias := c.Params("alias")
	identity := models.Identity{}
	result := database.DB.Db.Where("alias = ?", alias).First(&identity)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	p, _ := crypto.HexToECDSA(identity.ECDSA)
	web3.GenerateBlockchainAddressFile(identity.Alias, crypto.PubkeyToAddress(p.PublicKey))

	return c.Status(200).JSON(identity)
}
