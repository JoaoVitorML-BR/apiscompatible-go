package server

import (
    "fmt"
    "login/application/routes"
    "login/infra/tools/database/config"
    "net/http"

    "github.com/gorilla/handlers"
)

func Server() {
    config.LoadInfos()

    fmt.Printf(":%d", config.Port)

    fmt.Println("Server running at port... localhost:", config.Port)
    http.ListenAndServe(fmt.Sprintf(":%d", config.Port), handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}),
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
        handlers.AllowCredentials(),
    )(routes.Routes()))
}
