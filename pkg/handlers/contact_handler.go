package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/trulyworthless/chatter-blocks/pkg/crypt"
	"github.com/trulyworthless/chatter-blocks/pkg/database"
	"github.com/trulyworthless/chatter-blocks/pkg/database/models"
)

func CreateContact(c *fiber.Ctx) error {
	contact := new(models.Contact)
	if err := c.BodyParser(contact); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	fmt.Println(contact.Alias)
	contact.PublicKey = crypt.RetrieveRSAPublicKeyFromFile(contact.Alias)
	database.DB.Db.Create(&contact)

	return c.Status(200).JSON(contact)
}

func GetContacts(c *fiber.Ctx) error {
	contacts := []models.Contact{}
	result := database.DB.Db.Find(&contacts)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(contacts)
}

func GetContactById(c *fiber.Ctx) error {
	id := c.Params("id")
	contacts := models.Contact{}
	result := database.DB.Db.First(&contacts, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(contacts)
}

func GetContactByAlias(c *fiber.Ctx) error {
	alias := c.Params("alias")
	contacts := models.Contact{}
	result := database.DB.Db.First(&contacts, alias)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(contacts)
}

// TODO expand upon updates
func UpdateContactById(c *fiber.Ctx) error {
	id := c.Params("id")
	contact := new(models.Contact)

	if err := c.BodyParser(contact); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.DB.Db.Where("id = ?", id).Updates(&contact)
	return c.Status(200).JSON(contact)
}

func UpdateContactByAlias(c *fiber.Ctx) error {
	alias := c.Params("alias")
	contact := new(models.Contact)

	if err := c.BodyParser(contact); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.DB.Db.Where("alias = ?", alias).Updates(&contact)
	return c.Status(200).JSON(contact)
}

// TODO remove WHERE "contacts"."deleted_at" IS NULL
func DeleteContactById(c *fiber.Ctx) error {
	id := c.Params("id")
	contact := models.Contact{}
	result := database.DB.Db.Where("id = ?", id).Delete(&contact)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(contact)
}

func DeleteContactByAlias(c *fiber.Ctx) error {
	alias := c.Params("alias")
	contact := models.Contact{}
	result := database.DB.Db.Where("alias = ?", alias).Delete(&contact)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(contact)
}
