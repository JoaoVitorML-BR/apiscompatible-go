package sucess

import "net/http"

func SuccessCreateUser(w http.ResponseWriter){
	w.Write([]byte("User created successfully"))
}