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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "contact not created: bad input", "data": err})
	}

	contact.PublicKey = crypt.RetrieveRSAPublicKeyFromFile(contact.Correspondent)
	contact.Address = web3.RetrieveBlockchainAddressFromFile(contact.Correspondent).Hex()

	if err := database.Db.Create(&contact).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "contact not created", "data": err})
	}

	return c.Status(fiber.StatusCreated).JSON(contact)
}

func GetContacts(c *fiber.Ctx) error {
	contacts := []models.Contact{}
	result := database.Db.Find(&contacts)

	if result.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(contacts)
}

func GetContactByCorrespondent(c *fiber.Ctx) error {
	correspondent := c.Params("correspondent")
	contact := models.Contact{}
	result := database.Db.Where("correspondent = ?", correspondent).First(&contact)

	if result.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(contact)
}

// TODO expand upon updates
// TODO include patches
func UpdateContactByCorrespondent(c *fiber.Ctx) error {
	correspondent := c.Params("correspondent")
	contact := new(models.Contact)

	if err := c.BodyParser(contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "contact not updated: bad input", "data": err})
	}

	contact.PublicKey = crypt.RetrieveRSAPublicKeyFromFile(contact.Correspondent)
	contact.Address = web3.RetrieveBlockchainAddressFromFile(contact.Correspondent).Hex()

	if err := database.Db.Where("correspondent = ?", correspondent).Updates(&contact).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "identity not updated", "data": err})
	}

	return c.Status(fiber.StatusOK).JSON(contact)
}

// TODO remove WHERE "contacts"."deleted_at" IS NULL
func DeleteContactByCorrespondent(c *fiber.Ctx) error {
	correspondent := c.Params("correspondent")
	contact := models.Contact{}

	if err := database.Db.Where("correspondent = ?", correspondent).Delete(&contact).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "contact not deleted", "data": err})
	}

	return c.Status(fiber.StatusNoContent).JSON(contact)
}
