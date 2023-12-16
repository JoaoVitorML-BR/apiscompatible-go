package controllers

import (
	"encoding/json"
	"io/ioutil"
	"login/cmd/handlers/utils"
	"login/cmd/handlers/utils/erruser"
	"login/cmd/handlers/utils/responses"
	"login/domain/secure"
	"login/domain/validation"
	"login/infra/mysql/bridge"
	"login/infra/tools/database"
	"login/infra/tools/database/errDB"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// GetUsers gets all users.
// @Summary Get all users
// @Description Gets a list of all users
// @ID get-all-users
// @Produce json
// @Success 200 {array} User
// @Failure 400 {string} string "Bad Request"
// @Router /user [get]
func CreatUser(w http.ResponseWriter, r *http.Request) {
	user, err := handlers.ConvertUserJsonToObject(w, r)
	if err != nil {
		erruser.ErrMessageConvertUserJsonToObject(w, err)
		return
	}

	validationUser := &validation.ValidStruct{UserParams: &bridge.CreateUserParams{
		Name:     user.Name,
		Password: string(user.Password),
	}}

	if err := validationUser.Prepare("createNewUser"); err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.DatabaseAPI()
	if err != nil {
		errDB.ErrConnectDB(w, err)
		return
	}
	defer db.Close()

	newQuerie := bridge.New(db)

	// Hash the password before creating the user**
	hashedPassword, err := secure.Hash(string(user.Password))
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	createUserParams := bridge.CreateUserParams{
		Name:     user.Name,
		Password: string(hashedPassword),
	}

	if err := newQuerie.CreateUser(r.Context(), createUserParams); err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.Json(w, http.StatusCreated, nil)
}

// GetUsers gets all users.
// @Summary Get all users
// @Description Gets a list of all users
// @ID get-all-users
// @Produce json
// @Success 200 {array} User
// @Failure 400 {string} string "Bad Request"
// @Router /user [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.DatabaseAPI()

	if err != nil {
		errDB.ErrConnectDB(w, err)
		return
	}

	defer db.Close()

	bridge := bridge.New(db)

	userData, err := bridge.FindUsers(r.Context())
	if err != nil {
		erruser.ErrMessageSearchUser(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(userData); err != nil {
		erruser.ErrMessageConvertUserToJson(w, err)
		return
	}
}

// GetUserById gets a user by ID.
// @Summary Get user by ID
// @Description Gets a user by ID
// @ID get-user-by-id
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Router /user/{id} [get]
func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := handlers.ConvertID(w, params["id"])
	if err != nil {
		erruser.ErrMessageConvertIDtoInt(w, err)
		return
	}

	db, err := database.DatabaseAPI()
	if err != nil {
		errDB.ErrConnectDB(w, err)
		return
	}
	defer db.Close()

	bridge := bridge.New(db)

	userData, err := bridge.FindUserByID(r.Context(), int32(ID))
	if err != nil {
		erruser.ErrMessageSearchUser(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(userData); err != nil {
		erruser.ErrMessageConvertUserToJson(w, err)
		return
	}
}

func GetUsersByName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	nameUser := strings.ToLower(params["user_name"])
	db, err := database.DatabaseAPI()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	bridge := bridge.New(db)

	userData, err := bridge.FindUserByName(r.Context(), "%"+nameUser+"%") // % helps to search if the name contains the value provided regardless of the location
	if err != nil {
		erruser.ErrMessageConvertUserToJson(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(userData); err != nil {
		erruser.ErrMessageConvertUserToJson(w, err)
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := handlers.ConvertID(w, params["id"])
	if err != nil {
		erruser.ErrMessageConvertIDtoInt(w, err)
		return
	}

	bodyReq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user bridge.UpdateUserParams
	if err = json.Unmarshal(bodyReq, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	validationUser := &validation.ValidStruct{UserParams: &bridge.UpdateUserParams{
		Name: user.Name,
		ID:   user.ID,
	}}
	if err := validationUser.Prepare("edit"); err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := validationUser.Prepare("edit"); err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.DatabaseAPI()
	if err != nil {
		errDB.ErrConnectDB(w, err)
		return
	}
	defer db.Close()

	newQuerie := bridge.New(db)

	if err = newQuerie.UpdateUser(r.Context(), bridge.UpdateUserParams{
		Name: user.Name,
		ID:   int32(ID),
	}); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.Json(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := handlers.ConvertID(w, params["id"])
	if err != nil {
		erruser.ErrMessageConvertIDtoInt(w, err)
		return
	}

	db, err := database.DatabaseAPI()
	if err != nil {
		errDB.ErrConnectDB(w, err)
		return
	}
	defer db.Close()

	bridge := bridge.New(db)

	err = bridge.DeleteUser(r.Context(), int32(ID))
	if err != nil {
		erruser.ErrMessageSearchUser(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	responses.Json(w, http.StatusOK, map[string]string{"message": "Usuário deletado com sucesso"})
}