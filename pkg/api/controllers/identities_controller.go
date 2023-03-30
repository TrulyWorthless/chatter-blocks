package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/trulyworthless/chatter-blocks/pkg/api/auth"
	"github.com/trulyworthless/chatter-blocks/pkg/api/models"
	"github.com/trulyworthless/chatter-blocks/pkg/api/responses"
	"github.com/trulyworthless/chatter-blocks/pkg/api/util/formaterror"
)

func (server *Server) CreateIdentity(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	identity := models.Identity{}
	err = json.Unmarshal(body, &identity)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	identity.Init()
	err = identity.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	identityCreated, err := identity.SaveIdentity(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, identityCreated.Uuid))
	responses.JSON(w, http.StatusCreated, identityCreated)
}

func (server *Server) GetIdentities(w http.ResponseWriter, r *http.Request) {
	identity := models.Identity{}
	identities, err := identity.FindAllIdentities(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, identities)
}

func (server *Server) GetIdentity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid, err := strconv.ParseUint(vars["uuid"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	identity := models.Identity{}
	identityGotten, err := identity.FindIdentityByID(server.DB, uint32(uuid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, identityGotten)
}

func (server *Server) UpdateIdentity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid, err := strconv.ParseUint(vars["uuid"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

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

	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	if tokenID != uint32(uuid) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	identity.Init()
	err = identity.Validate("update")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	updatedIdentity, err := identity.UpdateAnIdentity(server.DB, uint32(uuid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	responses.JSON(w, http.StatusOK, updatedIdentity)
}

func (server *Server) DeletIdentity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	identity := models.Identity{}
	uuid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	if tokenID != 0 && tokenID != uint32(uuid) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	_, err = identity.DeleteAnIdentity(server.DB, uint32(uuid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d", uuid))
	responses.JSON(w, http.StatusNoContent, "")
}
