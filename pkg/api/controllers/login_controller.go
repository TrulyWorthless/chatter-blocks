package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/trulyworthless/chatter-blocks/pkg/api/auth"
	"github.com/trulyworthless/chatter-blocks/pkg/api/models"
	"github.com/trulyworthless/chatter-blocks/pkg/api/responses"
	"github.com/trulyworthless/chatter-blocks/pkg/api/util/formaterror"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	identity := models.Identity{}
	err = json.Unmarshal(body, &identity)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	identity.Init()
	err = identity.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := server.SignIn(identity.Alias, identity.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(alias, password string) (string, error) {
	identity := models.Identity{}

	err := server.DB.Debug().Model(models.Identity{}).Where("alias = ?", alias).Take(&identity).Error
	if err != nil {
		return "", err
	}

	err = models.VerifyPassword(identity.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	return auth.CreateToken(identity.Uuid)
}
