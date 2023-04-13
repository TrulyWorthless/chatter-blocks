package handlers

import (
	"errors"
	"time"

	"github.com/trulyworthless/chatter-blocks/pkg/config"
	"github.com/trulyworthless/chatter-blocks/pkg/database"
	"github.com/trulyworthless/chatter-blocks/pkg/database/models"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func validToken(t *jwt.Token, username string) bool {
	claims := t.Claims.(jwt.MapClaims)
	uid := claims["username"]

	return uid == username
}

func Login(c *fiber.Ctx) error {
	var input UserData
	var ud UserData

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "cannot login: bad data", "data": err})
	}

	user, err := new(models.User), *new(error)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "cannot login: database error", "data": err})
	}

	if err := database.Db.Where("username = ?", input.Username).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "user not found", "data": err})
		} else {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
	}

	ud = UserData{
		Username: user.Username,
		Password: user.Password,
	}

	if !checkPasswordHash(input.Password, ud.Password) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = ud.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "login successful", "data": t})
}

func Logout(c *fiber.Ctx) error {
	var input UserData
	var ud UserData

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "cannot login: bad data", "data": err})
	}

	user, err := new(models.User), *new(error)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "cannot login: database error", "data": err})
	}

	if err := database.Db.Where("username = ?", input.Username).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "user not found", "data": err})
		} else {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
	}

	ud = UserData{
		Username: user.Username,
		Password: user.Password,
	}

	if !checkPasswordHash(input.Password, ud.Password) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, ud.Username) {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Unix()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "logout successful", "data": nil})
}
