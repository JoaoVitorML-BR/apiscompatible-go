package errDB

import (
	"log"
	"net/http"
)

func ErrConnectDB(w http.ResponseWriter, err error){
	w.Write([]byte("Error connecting to the database"))

	log.Printf("Error connecting to the database: %v", err)
}

