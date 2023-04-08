package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trulyworthless/chatter-blocks/pkg/crypt"
	"github.com/trulyworthless/chatter-blocks/pkg/database"
	"github.com/trulyworthless/chatter-blocks/pkg/database/models"
	"github.com/trulyworthless/chatter-blocks/pkg/web3"
)

func CreateContact(c *fiber.Ctx) error {
	contact := new(models.Contact)
	if err := c.BodyParser(contact); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	contact.PublicKey = crypt.RetrieveRSAPublicKeyFromFile(contact.Correspondent)
	contact.Address = web3.RetrieveBlockchainAddressFromFile(contact.Correspondent).Hex()
	database.DB.Db.Create(&contact)

	return c.Status(fiber.StatusCreated).JSON(contact)
}

func GetContacts(c *fiber.Ctx) error {
	contacts := []models.Contact{}
	result := database.DB.Db.Find(&contacts)

	if result.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(contacts)
}

func GetContactByCorrespondent(c *fiber.Ctx) error {
	correspondent := c.Params("correspondent")
	contact := models.Contact{}
	result := database.DB.Db.Where("correspondent = ?", correspondent).First(&contact)

	if result.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(contact)
}

// TODO expand upon updates
func UpdateContactByCorrespondent(c *fiber.Ctx) error {
	correspondent := c.Params("correspondent")
	contact := new(models.Contact)

	if err := c.BodyParser(contact); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	contact.PublicKey = crypt.RetrieveRSAPublicKeyFromFile(contact.Correspondent)
	contact.Address = web3.RetrieveBlockchainAddressFromFile(contact.Correspondent).Hex()

	database.DB.Db.Where("correspondent = ?", correspondent).Updates(&contact)
	return c.Status(fiber.StatusOK).JSON(contact)
}

// TODO remove WHERE "contacts"."deleted_at" IS NULL
func DeleteContactByCorrespondent(c *fiber.Ctx) error {
	correspondent := c.Params("correspondent")
	contact := models.Contact{}
	result := database.DB.Db.Where("correspondent = ?", correspondent).Delete(&contact)

	if result.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusNoContent).JSON(contact)
}
