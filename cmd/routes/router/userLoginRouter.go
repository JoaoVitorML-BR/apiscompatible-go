package router

import (
	"login/cmd/controllers"
	"net/http"
)

var RouterUsersLogin = Router{
    URI:         "/user/login",
    MethodRouter: http.MethodPost,
    FuncRouter:   controllers.Login,
    AuthRouter:   false,
}