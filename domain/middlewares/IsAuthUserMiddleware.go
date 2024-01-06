package middlewares

import (
	"log"
	"login/application/handlers/utils/responses"
	"login/domain/auth"
	"net/http"
)

// Logger writer info about req on terminal
func Logger(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Checks if the user doing the authentication is authenticated.
func IsAuthUser(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		if err := auth.ValidToken(r); err != nil {
			responses.Err(w, http.StatusUnauthorized, err);
			return
		}
		next(w, r)
	}
}