package handlers

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gofiber/fiber/v2"
	"github.com/trulyworthless/chatter-blocks/pkg/crypt"
	"github.com/trulyworthless/chatter-blocks/pkg/database"
	"github.com/trulyworthless/chatter-blocks/pkg/database/models"
	"github.com/trulyworthless/chatter-blocks/pkg/web3"
)

// TODO: combine files
func ExportPublicKeyById(c *fiber.Ctx) error {
	id := c.Params("id")
	identity := models.Identity{}
	result := database.Db.First(&identity, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	err := crypt.GenerateRSAPublicKeyFile(identity.Alias, &crypt.GetRSAPrivateKeyFromBytes(identity.RSA).PublicKey)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "identity publickey not exported", "data": err})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "identity publickey exported", "data": identity.Alias})
}

func ExportPublicKeyByAlias(c *fiber.Ctx) error {
	alias := c.Params("alias")
	identity := models.Identity{}
	result := database.Db.Where("alias = ?", alias).First(&identity)

	if result.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	err := crypt.GenerateRSAPublicKeyFile(identity.Alias, &crypt.GetRSAPrivateKeyFromBytes(identity.RSA).PublicKey)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "identity publickey not exported", "data": err})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "identity publickey exported", "data": identity.Alias})
}

// TODO use common.Address
func ExportAddressById(c *fiber.Ctx) error {
	id := c.Params("id")
	identity := models.Identity{}
	result := database.Db.First(&identity, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	p, _ := crypto.HexToECDSA(identity.ECDSA)
	err := web3.GenerateBlockchainAddressFile(identity.Alias, crypto.PubkeyToAddress(p.PublicKey))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "identity address not exported", "data": err})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "identity address exported", "data": identity.Alias})
}

func ExportAddressByAlias(c *fiber.Ctx) error {
	alias := c.Params("alias")
	identity := models.Identity{}
	result := database.Db.Where("alias = ?", alias).First(&identity)

	if result.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	p, _ := crypto.HexToECDSA(identity.ECDSA)
	err := web3.GenerateBlockchainAddressFile(identity.Alias, crypto.PubkeyToAddress(p.PublicKey))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "identity address not exported", "data": err})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "identity address exported", "data": identity.Alias})
}
