package erruser

import (
	"encoding/json"
	"log"
	"net/http"
)

func ErrMessageCreatUserJson(w http.ResponseWriter, statusCode int, data interface{}){
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil{
		log.Fatal(err)
	}
}

func ErrMessageCreateUser(w http.ResponseWriter, statusCode int, err error){
	ErrMessageCreatUserJson(w, statusCode, struct{
		Err string `json:"erro"`
	}{
		Err: err.Error(),
	})
}