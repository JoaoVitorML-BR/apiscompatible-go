package auth

import (
	"login/infra/tools/database/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenToken(userID uint64) (string, error){
	// permissions for token
	permissions := jwt.MapClaims{}
	permissions["authoried"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 10).Unix()
	permissions["userID"] = userID

	//Generate the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey)) // sign the token
}