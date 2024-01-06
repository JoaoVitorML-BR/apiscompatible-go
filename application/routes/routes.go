package routes

import (
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	r := mux.NewRouter();

	return ConfigRouter(r)
}
