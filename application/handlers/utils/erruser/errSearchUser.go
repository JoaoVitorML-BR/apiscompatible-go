package erruser

import (
	"log"
	"net/http"
)

func ErrMessageSearchUser(w http.ResponseWriter, err error){
	w.Write([]byte("Error when searching for 'user' by 'id'"))

	log.Printf("Error when searching for 'user' by 'id'")
}