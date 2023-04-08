package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trulyworthless/chatter-blocks/pkg/crypt"
	"github.com/trulyworthless/chatter-blocks/pkg/database"
	"github.com/trulyworthless/chatter-blocks/pkg/database/models"
)

// TODO hide private key
func CreateIdentity(c *fiber.Ctx) error {
	identity := new(models.Identity)
	identity.RSA = crypt.GenerateRSAPrivateKeyToBytes()
	if err := c.BodyParser(identity); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "identity not created: bad input", "data": err})
	}

	if err := database.DB.Db.Create(&identity).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "identity not created", "data": err})
	}

	//TODO
	return c.Status(fiber.StatusCreated).JSON(identity)
}

func GetIdentities(c *fiber.Ctx) error {
	identities := []models.Identity{}
	result := database.DB.Db.Find(&identities)

	if result.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(identities)
}

func GetIdentityByAlias(c *fiber.Ctx) error {
	alias := c.Params("alias")
	identity := models.Identity{}
	result := database.DB.Db.Where("alias = ?", alias).First(&identity)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "alias not found", "data": nil})
	}

	return c.Status(fiber.StatusOK).JSON(identity)
}

// TODO expand upon updates
func UpdateIdentityByAlias(c *fiber.Ctx) error {
	alias := c.Params("alias")
	identity := new(models.Identity)

	if err := c.BodyParser(identity); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "identity not updated: bad input", "data": err})
	}

	database.DB.Db.Where("alias = ?", alias).Updates(&identity)
	return c.Status(fiber.StatusOK).JSON(identity)
}

// TODO remove WHERE "contacts"."deleted_at" IS NULL
func DeleteIdentityByAlias(c *fiber.Ctx) error {
	alias := c.Params("alias")
	identity := models.Identity{}
	result := database.DB.Db.Where("alias = ?", alias).Delete(&identity)

	if result.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusNoContent).JSON(identity)
}
