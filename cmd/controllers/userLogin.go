package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"login/cmd/handlers/utils/erruser"
	"login/cmd/handlers/utils/responses"
	"login/domain/auth"
	"login/domain/secure"
	"login/infra/mysql/bridge"
	"login/infra/tools/database"
	"login/infra/tools/database/errDB"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	bodyReq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user bridge.FindUserByNameToLoginRow
	if err = json.Unmarshal(bodyReq, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.DatabaseAPI()
	if err != nil {
		errDB.ErrConnectDB(w, err)
		return
	}
	defer db.Close()

	newQuerie := bridge.New(db)

	searchUserOnDB, err := newQuerie.FindUserByNameToLogin(r.Context(), user.Name)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	isEqual := secure.ComparePasswordWithHash(string(searchUserOnDB.Password), string(user.Password))
	if isEqual == nil {
		fmt.Print("As senhas são iguais \n")
	} else {
		fmt.Print("Senha incorreta\n")
		responses.Err(w, http.StatusUnauthorized, errors.New("Senha incorreta"))
		return
	}

	token, _ := auth.GenToken(uint64(user.ID))

	fmt.Print(token)

	w.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(w).Encode(searchUserOnDB); err != nil {
		erruser.ErrMessageConvertUserToJson(w, err)
		return
	}
}
