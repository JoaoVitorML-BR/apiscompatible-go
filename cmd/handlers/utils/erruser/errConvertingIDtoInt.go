package erruser

import (
	"log"
	"net/http"
)

func ErrMessageConvertIDtoInt(w http.ResponseWriter, err error){
	w.Write([]byte("Error to converting ID user"))

	log.Printf("Error to converting ID user: %v", err)
}