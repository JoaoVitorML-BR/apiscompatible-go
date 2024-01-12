package auth

import (
	"errors"
	"fmt"
	"login/infra/tools/database/config"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Create token
func GenToken(userID uint64) (string, error) {
	// permissions for token
	permissions := jwt.MapClaims{}
	permissions["authoried"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 10).Unix()
	permissions["id"] = userID 

	//Generate the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	fmt.Print(token)
	return token.SignedString([]byte(config.SecretKey)) // sign the token
}

// extract token that's Header browser "Authorization" (example in insominia (Auth))
func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func ExtractUserID(r *http.Request) (uint64, error) {
	tokenString := extractToken(r) // get token extract
	token, err := jwt.Parse(tokenString, returnKeyCheckToekn)
	if err != nil {
		return 0, err
	}

	if permission, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", permission["id"]), 10, 64) // fix to "id" (Uint)
		if err != nil {
			return 0, err
		}

		return userID, nil
	}

	return 0, errors.New("Token Invalido!")
}

// checks if the passed token is valid
func ValidToken(r *http.Request) error {
	tokenString := extractToken(r) // get token extract
	token, err := jwt.Parse(tokenString, returnKeyCheckToekn)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token inválido!")
}

func ValidateTokenHandler(w http.ResponseWriter, r *http.Request) {
    err := ValidToken(r)
    if err != nil {
        http.Error(w, "Token inválido", http.StatusUnauthorized)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Token válido"))
}

func returnKeyCheckToekn(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Signature method is not correct", token.Header["alg"])
	}

	return config.SecretKey, nil
}
