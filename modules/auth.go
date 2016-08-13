package modules

import (
	"database/sql"
	"net/http"
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"

	"reactizer-go/config"
)

type AuthError string

func (e AuthError) Error() string {
	return string(e)
}

// 'authorize' checks the 'X-Authorization' header if it contains the JWT token required by some
// queries. If the token is there, it is decoded into a user id and returned.
//
// In case of an error, translation id AuthError is returned.
func authorize(r *http.Request, db *sql.DB) (int, error) {
	data := r.Header["X-Authorization"]
	if len(data) != 1 {
		return 0, AuthError("auth.no_auth_header")
	}
	token := data[0]

	log.Print(decodeToken(token)) // TODO: create token
	return 0, nil
}

func decodeToken(raw string) (int, error) {
	token, err := jwt.Parse(raw, keyfunc)
	if err != nil {
		log.Print(err)
		return 0, AuthError("auth.invalid_token")
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    //return claims["sub"].(int), nil
    return 1, nil
	}

	return 0, AuthError("auth.invalid_token")
}

func keyfunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected method: %v", token.Header["alg"])
	}
	return config.Secret, nil
}
