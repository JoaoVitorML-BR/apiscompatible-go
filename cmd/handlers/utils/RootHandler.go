package handlers

import (
	"login/infra/tools/database"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	// Trate a solicitação para a raiz aqui
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Página de raiz"))

	database.DatabaseAPI()
}
