package routes

import (
	// "fmt"
	"fmt"
	"login/cmd/routes/router"
	// "login/cmd/tools/database/config"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
)

type Router struct {
    URI         string
    MethodRouter string
    FuncRouter   func(w http.ResponseWriter, r *http.Request)
    AuthRouter   bool
}

func ConfigRouter(r *mux.Router) *mux.Router {
    RoutesUsersConfig := router.RoutersUsers
    RoutesUsersConfig = append(RoutesUsersConfig, router.RouterUsersLogin)


    for _, route := range RoutesUsersConfig {
        r.HandleFunc(route.URI, route.FuncRouter).Methods(route.MethodRouter)
    }

    r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
        httpSwagger.URL(fmt.Sprintf("/swagger/doc.json"), // certifique-se de que o caminho seja correto
    )))

    return r
}
