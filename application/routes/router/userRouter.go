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
        URI:           "/go-api/valid-token",
        MethodRouter:  http.MethodGet,
        FuncRouter:    auth.ValidateTokenHandler,
        AuthRouter:    true,
    },
	// User router
	{
		URI: "/go-api/user",
		MethodRouter: http.MethodPost,
		FuncRouter: controllers.CreatUser,
		AuthRouter: true,
	},
	{
		URI: "/go-api/user",
		MethodRouter: http.MethodGet,
		FuncRouter: controllers.GetUsers,
		AuthRouter: true,
	},
	{
		URI: "/go-api/user/{id}",
		MethodRouter: http.MethodGet,
		FuncRouter: controllers.GetUserById,
		AuthRouter: true,
	},{
		URI: "/go-api/userName/{user_name}",
		MethodRouter: http.MethodGet,
		FuncRouter: controllers.GetUsersByName,
		AuthRouter: true,
	},
	{
		URI: "/go-api/user/{id}",
		MethodRouter: http.MethodPut,
		FuncRouter: controllers.UpdateUser,
		AuthRouter: true,
	},{
		URI: "/go-api/user/{id}",
		MethodRouter: http.MethodDelete,
		FuncRouter: controllers.DeleteUser,
		AuthRouter: true,
	},
}
