package erruser

import (
	"log"
	"net/http"
)

func ErrMessageConvertUserToJson(w http.ResponseWriter, err error){
	w.Write([]byte("Error passing 'user' to 'json'"))

	log.Printf("Error passing 'user' to 'json'")
}