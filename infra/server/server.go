package server

import (
	"fmt"
	"login/application/routes"
	"login/infra/tools/database/config"
	"net/http"
)

func Server() {
	config.LoadInfos()

	fmt.Printf(":%d", config.Port)

	fmt.Println("Server running at port... localhost:", config.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), routes.Routes())
}
