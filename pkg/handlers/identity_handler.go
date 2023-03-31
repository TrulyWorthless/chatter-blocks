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
	if err := c.BodyParser(identity); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.DB.Db.Create(&identity)

	//TODO insecury, do not send back private key
	return c.Status(200).JSON(identity)
}

func GetIdentities(c *fiber.Ctx) error {
	identities := []models.Identity{}
	result := database.DB.Db.Find(&identities)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(identities)
}

func GetIdentityById(c *fiber.Ctx) error {
	id := c.Params("id")
	identity := models.Identity{}
	result := database.DB.Db.First(&identity, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(identity)
}

func GetIdentityByAlias(c *fiber.Ctx) error {
	alias := c.Params("alias")
	identity := models.Identity{}
	result := database.DB.Db.Where("alias = ?", alias).First(&identity)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(identity)
}

// TODO expand upon updates
func UpdateIdentityByID(c *fiber.Ctx) error {
	id := c.Params("id")
	identity := new(models.Identity)

	if err := c.BodyParser(identity); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.DB.Db.Where("id = ?", id).Updates(&identity)
	return c.Status(200).JSON(identity)
}

func UpdateIdentityByAlias(c *fiber.Ctx) error {
	alias := c.Params("alias")
	identity := new(models.Identity)

	if err := c.BodyParser(identity); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.DB.Db.Where("alias = ?", alias).Updates(&identity)
	return c.Status(200).JSON(identity)
}

func DeletIdentityByID(c *fiber.Ctx) error {
	id := c.Params("id")
	identity := models.Identity{}
	result := database.DB.Db.Where("id = ?", id).Delete(&identity)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(identity)
}

func DeletIdentityByAlias(c *fiber.Ctx) error {
	alias := c.Params("alias")
	identity := models.Identity{}
	result := database.DB.Db.Where("alias = ?", alias).Delete(&identity)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(identity)
}
