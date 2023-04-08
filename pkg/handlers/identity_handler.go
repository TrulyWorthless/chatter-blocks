package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trulyworthless/chatter-blocks/pkg/crypt"
	"github.com/trulyworthless/chatter-blocks/pkg/database"
	"github.com/trulyworthless/chatter-blocks/pkg/database/models"
)

func CreateIdentity(c *fiber.Ctx) error {
	identity := new(models.Identity)
	identity.RSA = crypt.GenerateRSAPrivateKeyToBytes()
	//TODO check if identity has edcas
	if err := c.BodyParser(identity); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "identity not created: bad input", "data": err})
	}

	if err := database.Db.Create(&identity).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "identity not created", "data": err})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "identity found", "data": identity.Alias})
}

func GetIdentities(c *fiber.Ctx) error {
	identities := []models.Identity{}
	result := database.Db.Find(&identities)

	if result.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	// TODO hide private key
	return c.Status(fiber.StatusOK).JSON(identities)
}

func GetIdentityByAlias(c *fiber.Ctx) error {
	alias := c.Params("alias")
	identity := models.Identity{}
	result := database.Db.Where("alias = ?", alias).First(&identity)

	if result.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(identity.Alias)
}

// TODO expand upon updates
// TODO include patches
func UpdateIdentityByAlias(c *fiber.Ctx) error {
	alias := c.Params("alias")
	identity := new(models.Identity)

	if err := c.BodyParser(identity); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "identity not updated: bad input", "data": err})
	}

	if err := database.Db.Where("alias = ?", alias).Updates(&identity).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "identity not updated", "data": err})
	}

	return c.Status(fiber.StatusOK).JSON(identity.Alias)
}

func DeleteIdentityByAlias(c *fiber.Ctx) error {
	alias := c.Params("alias")
	identity := models.Identity{}

	if err := database.Db.Where("alias = ?", alias).Delete(&identity).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "identity not deleted", "data": err})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
