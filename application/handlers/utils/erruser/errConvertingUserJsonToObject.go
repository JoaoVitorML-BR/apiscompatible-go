package erruser

import (
	"log"
	"net/http"
)

func ErrMessageConvertUserJsonToObject(w http.ResponseWriter, err error){
	w.Write([]byte("Error to converting User(json) to User(Object type Struct)"))

	log.Printf("Error to converting User(json) to User(Object type Struct): %v", err)
}