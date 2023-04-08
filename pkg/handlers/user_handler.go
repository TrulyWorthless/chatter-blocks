package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/trulyworthless/chatter-blocks/pkg/database"
	"github.com/trulyworthless/chatter-blocks/pkg/database/models"
)

type NewUser struct {
	Username string `json:"username"`
}

type PasswordInput struct {
	Password    string `json:"password"`
	NewPassword string `json:"newpass"`
}

func validToken(t *jwt.Token, username string) bool {
	claims := t.Claims.(jwt.MapClaims)
	uid := claims["username"]

	return uid == username
}

func validUser(username string, p string) (models.User, bool) {
	user := models.User{}
	result := database.DB.Db.Where("username = ?", username).First(&user)
	if result.RowsAffected == 0 {
		return models.User{}, false
	}

	if !CheckPasswordHash(p, user.Password) {
		return models.User{}, false
	}

	return user, true
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "user not created: bad input", "data": err})
	}

	hash, err := HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "user not created: invalid password", "data": err})
	}

	user.Password = hash
	if err := database.DB.Db.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "user not created", "data": err})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "user created", "data": user.Username})
}

func GetUser(c *fiber.Ctx) error {
	username := c.Params("username")
	user := models.User{}
	result := database.DB.Db.Where("username = ?", username).First(&user)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "user not found", "data": nil})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "user found", "data": user.Username})
}

func UpdateUser(c *fiber.Ctx) error {
	var pi PasswordInput
	if err := c.BodyParser(&pi); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "user not updated: bad input", "data": err})
	}

	username := c.Params("username")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, username) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "user not updated: invalid token id", "data": nil})
	}

	user, valid := validUser(username, pi.Password)
	if !valid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "user not updated: user information not valid", "data": nil})
	}

	user.Password, _ = HashPassword(pi.NewPassword)
	if err := database.DB.Db.Where("username = ?", username).Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "user not updated", "data": err})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "user updated", "data": user.Username})
}

func DeleteUser(c *fiber.Ctx) error {
	var pi PasswordInput
	if err := c.BodyParser(&pi); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "user not deleted: bad input", "data": err})
	}

	username := c.Params("username")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, username) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "user not deleted: invalid token id", "data": nil})
	}

	user, valid := validUser(username, pi.Password)
	if !valid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "user not deleted: user information not valid", "data": nil})
	}

	if err := database.DB.Db.Where("username = ?", username).Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "user not deleted", "data": err})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"status": "success", "message": "user deleted", "data": nil})
}
