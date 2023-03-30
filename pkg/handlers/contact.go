package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func CreateContact(c *fiber.Ctx) error {
	// body, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }

	// contact := models.Contact{}
	// err = json.Unmarshal(body, &contact)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }

	// contact.Init()
	// err = contact.Validate()
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }

	// uuid, err := auth.ExtractTokenID(r)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }

	// if uuid != contact.Uuid {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
	// 	return
	// }
	// contactCreated, err := contact.SaveContact(server.DB)
	// if err != nil {
	// 	formattedError := formaterror.FormatError(err.Error())
	// 	responses.ERROR(w, http.StatusInternalServerError, formattedError)
	// 	return
	// }

	// w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, contactCreated.Uuid))
	// responses.JSON(w, http.StatusCreated, contactCreated)
	return errors.New("identities: CreateContact")
}

func GetContacts(c *fiber.Ctx) error {
	// contact := models.Contact{}
	// contacts, err := contact.FindAllContacts(server.DB)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusInternalServerError, err)
	// 	return
	// }
	// responses.JSON(w, http.StatusOK, contacts)
	return errors.New("identities: CreateCGetContactsontact")
}

func GetContact(c *fiber.Ctx) error {
	// vars := mux.Vars(r)
	// uuid, err := strconv.ParseUint(vars["uuid"], 10, 64)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusBadRequest, err)
	// 	return
	// }
	// contact := models.Contact{}

	// contactReceived, err := contact.FindContactByID(server.DB, uint32(uuid))
	// if err != nil {
	// 	responses.ERROR(w, http.StatusInternalServerError, err)
	// 	return
	// }
	// responses.JSON(w, http.StatusOK, contactReceived)
	return errors.New("identities: GetContact")
}

func UpdateContact(c *fiber.Ctx) error {
	// vars := mux.Vars(r)

	// // Check if the contact id is valid
	// uuid, err := strconv.ParseUint(vars["uuid"], 10, 64)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusBadRequest, err)
	// 	return
	// }

	// //CHeck if the auth token is valid and  get the user id from it
	// Uuid, err := auth.ExtractTokenID(r)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }

	// // Check if the contact exist
	// contact := models.Contact{}
	// err = server.DB.Debug().Model(models.Contact{}).Where("id = ?", uuid).Take(&contact).Error
	// if err != nil {
	// 	responses.ERROR(w, http.StatusNotFound, errors.New("contact not found"))
	// 	return
	// }

	// // If a user attempt to update a contact not belonging to him
	// if Uuid != contact.Uuid {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }
	// // Read the data contact
	// body, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }

	// // Start processing the request data
	// contactUpdate := models.Contact{}
	// err = json.Unmarshal(body, &contactUpdate)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }

	// //Also check if the request user id is equal to the one gotten from token
	// if Uuid != contactUpdate.Uuid {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }

	// contactUpdate.Init()
	// err = contactUpdate.Validate()
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }

	// contactUpdate.Uuid = contact.Uuid //this is important to tell the model the contract id to update, the other update field are set above

	// contactUpdated, err := contactUpdate.UpdateAContact(server.DB, contactUpdate.Uuid)

	// if err != nil {
	// 	formattedError := formaterror.FormatError(err.Error())
	// 	responses.ERROR(w, http.StatusInternalServerError, formattedError)
	// 	return
	// }
	// responses.JSON(w, http.StatusOK, contactUpdated)
	return errors.New("identities: UpdateContact")
}

func DeleteContact(c *fiber.Ctx) error {

	// vars := mux.Vars(r)

	// // Is a valid contact id given to us?
	// uuid, err := strconv.ParseUint(vars["uuid"], 10, 64)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusBadRequest, err)
	// 	return
	// }

	// // Is this user authenticated?
	// Uuid, err := auth.ExtractTokenID(r)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }

	// // Check if the contact exist
	// contact := models.Contact{}
	// err = server.DB.Debug().Model(models.Contact{}).Where("id = ?", uuid).Take(&contact).Error
	// if err != nil {
	// 	responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
	// 	return
	// }

	// // Is the authenticated user, the owner of this contact?
	// if Uuid != contact.Uuid {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }
	// _, err = contact.DeleteAContact(server.DB, Uuid)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusBadRequest, err)
	// 	return
	// }
	// w.Header().Set("Entity", fmt.Sprintf("%d", uuid))
	// responses.JSON(w, http.StatusNoContent, "")
	return errors.New("identities: DeleteContact")
}
