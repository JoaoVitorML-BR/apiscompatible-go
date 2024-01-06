package routes

import (
	"fmt"

	"login/application/routes/router"
	"login/domain/middlewares"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

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

	corsMiddleware := cors.New(cors.Options{
        AllowedOrigins:     []string{"*"},
        AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowedHeaders:     []string{"Content-Type", "Authorization"},
        AllowCredentials:   true,
    })
    
    r.Use(corsMiddleware.Handler)
    
    for _, route := range RoutesUsersConfig {
        var handler http.HandlerFunc
    
        if route.AuthRouter {
            handler = middlewares.Logger(middlewares.IsAuthUser(route.FuncRouter))
        } else {
            handler = middlewares.Logger(route.FuncRouter)
        }
    
        r.HandleFunc(route.URI, handler).Methods(route.MethodRouter)
    }

    r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
        httpSwagger.URL(fmt.Sprintf("/swagger/doc.json"), // certifique-se de que o caminho seja correto
    )))

    return r
}