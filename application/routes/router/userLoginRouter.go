package router

import (
	"login/application/controllers"
	"net/http"
)

var RouterUsersLogin = Router{
    URI:         "/login",
    MethodRouter: http.MethodPost,
    FuncRouter:   controllers.Login,
    AuthRouter:   false,
}