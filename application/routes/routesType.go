// routes.go
package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"login/application/routes/router"
	"login/domain/middlewares"
)

func ConfigRouter(r *mux.Router) *mux.Router {
	RoutesUsersConfig := router.RoutersUsers
	RoutesUsersConfig = append(RoutesUsersConfig, router.RouterUsersLogin)

	for _, route := range RoutesUsersConfig {
		var handler http.HandlerFunc

		if route.AuthRouter {
			handler = middlewares.Logger(middlewares.IsAuthUser(route.FuncRouter))
		} else {
			handler = middlewares.Logger(route.FuncRouter)
		}

		r.HandleFunc(route.URI, handler).Methods(route.MethodRouter)
	}

	return r
}