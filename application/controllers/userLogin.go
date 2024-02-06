package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"login/application/handlers/utils/erruser"
	"login/application/handlers/utils/responses"
	"login/domain/auth"
	"login/domain/secure"
	"login/infra/mysql/bridge"
	"login/infra/tools/database"
	"login/infra/tools/database/errDB"
	"net/http"
)
type LoginResponse struct {
	Token string               `json:"token"`
	User  bridge.FindUserByNameToLoginRow   `json:"user"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	bodyReq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		fmt.Println("Error reading request body:", err)
		return
	}

	fmt.Println(bodyReq)

	var user bridge.FindUserByNameToLoginRow
	if err = json.Unmarshal(bodyReq, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		fmt.Println("Error deserialize JSON:", err)
		return
	}

	db, err := database.DatabaseAPI()
	if err != nil {
		errDB.ErrConnectDB(w, err)
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	newQuerie := bridge.New(db)

	searchUserOnDB, err := newQuerie.FindUserByNameToLogin(r.Context(), user.Name)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		fmt.Println("Error when searching for user in the database:", err)
		return
	}

	isEqual := secure.ComparePasswordWithHash(string(searchUserOnDB.Password), string(user.Password))
	if isEqual == nil {
		fmt.Println("Password is equal")
	} else {
		fmt.Println("Incorrect password")
		responses.Err(w, http.StatusUnauthorized, errors.New("Incorrect password"))
		return
	}

	token, _ := auth.GenToken(uint64(searchUserOnDB.ID))

	fmt.Println("Token generated:", token)

	loginResponse := LoginResponse{
		Token: token,
		User:  searchUserOnDB,
	}

	w.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(w).Encode(loginResponse); err != nil {
		erruser.ErrMessageConvertUserToJson(w, err)
		fmt.Println("Error converting response to JSON:", err)
		return
	}
}