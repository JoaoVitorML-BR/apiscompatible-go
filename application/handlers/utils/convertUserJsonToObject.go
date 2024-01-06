package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"login/application/handlers/utils/erruser"
	"net/http"
)

func ConvertUserJsonToObject(w http.ResponseWriter,r *http.Request) (UserType, error){
	bodyReq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		erruser.ErrMessageConvertUserJsonToObject(w, err)
		fmt.Printf(string(bodyReq), UserType{})
		return UserType{}, err
	}

	fmt.Printf(string(bodyReq))

	var user UserType
	if err := json.Unmarshal(bodyReq, &user); err != nil {
		erruser.ErrMessageConvertUserJsonToObject(w, err)
		return UserType{}, err
	}

	return user, nil
}