package router

import (
	"login/application/controllers"
	"net/http"
)

var RouterUsersLogin = Router{
    URI:         "/go-api/login",
    MethodRouter: http.MethodPost,
    FuncRouter:   controllers.Login,
    AuthRouter:   false,
}