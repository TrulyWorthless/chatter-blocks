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

type LoginInput struct {
	Profile  string `json:"profile"`
	Password string `json:"password"`
}

type UserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(c *fiber.Ctx) error {
	input := new(LoginInput)
	var ud UserData

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "cannot login: bad data", "data": err})
	}

	profile := input.Profile
	password := input.Password
	user, err := new(models.User), *new(error)

	if err := database.Db.Where(&models.User{Username: profile}).Find(&user).Error; err != nil {
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

	if !CheckPasswordHash(password, ud.Password) {
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
