package router

import (
	"net/http"

	"login/cmd/controllers"
)

type Router struct {
	URI         string
	MethodRouter string
	FuncRouter   func(w http.ResponseWriter, r *http.Request)
	AuthRouter   bool
}

var RoutersUsers = []Router{
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
		AuthRouter: false,
	},{
		URI: "/userName/{user_name}",
		MethodRouter: http.MethodGet,
		FuncRouter: controllers.GetUsersByName,
		AuthRouter: false,
	},
	{
		URI: "/user/{id}",
		MethodRouter: http.MethodPut,
		FuncRouter: controllers.UpdateUser,
		AuthRouter: false,
	},{
		URI: "/user/{id}",
		MethodRouter: http.MethodDelete,
		FuncRouter: controllers.DeleteUser,
		AuthRouter: false,
	},
}
