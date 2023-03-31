package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

//TODO SignUp

func Login(c *fiber.Ctx) error {
	// body, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }

	// identity := models.Identity{}
	// err = json.Unmarshal(body, &identity)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }

	// identity.Init()
	// err = identity.Validate("login")
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }

	// token, err := server.SignIn(identity.Alias, identity.Password)
	// if err != nil {
	// 	formattedError := formaterror.FormatError(err.Error())
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
	// 	return
	// }

	// responses.JSON(w, http.StatusOK, token)
	return errors.New("identities: Login")
}

func SignIn(c *fiber.Ctx) error {
	// identity := models.Identity{}

	// err := server.DB.Debug().Model(models.Identity{}).Where("alias = ?", alias).Take(&identity).Error
	// if err != nil {
	// 	return "", err
	// }

	// err = models.VerifyPassword(identity.Password, password)
	// if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
	// 	return "", err
	// }

	// return auth.CreateToken(identity.Uuid)
	return errors.New("identities: SignIn")
}
