package router

import (
	"net/http"

	"login/domain/auth"
	"login/application/controllers"
)

type Router struct {
	URI         string
	MethodRouter string
	FuncRouter   func(w http.ResponseWriter, r *http.Request)
	AuthRouter   bool
}

var RoutersUsers = []Router{
	{
        URI:           "/valid-token",
        MethodRouter:  http.MethodGet,
        FuncRouter:    auth.ValidateTokenHandler,
        AuthRouter:    true,
    },
	// User router
	{
		URI: "/user",
		MethodRouter: http.MethodPost,
		FuncRouter: controllers.CreatUser,
		AuthRouter: false,
	},
	{
		URI: "/user",
		MethodRouter: http.MethodGet,
		FuncRouter: controllers.GetUsers,
		AuthRouter: false,
	},
	{
		URI: "/user/{id}",
		MethodRouter: http.MethodGet,
		FuncRouter: controllers.GetUserById,
		AuthRouter: true,
	},{
		URI: "/userName/{user_name}",
		MethodRouter: http.MethodGet,
		FuncRouter: controllers.GetUsersByName,
		AuthRouter: true,
	},
	{
		URI: "/user/{id}",
		MethodRouter: http.MethodPut,
		FuncRouter: controllers.UpdateUser,
		AuthRouter: true,
	},{
		URI: "/user/{id}",
		MethodRouter: http.MethodDelete,
		FuncRouter: controllers.DeleteUser,
		AuthRouter: true,
	},
}
