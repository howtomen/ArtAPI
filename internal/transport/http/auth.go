package http

import (
	"errors"
	"net/http"
	"strings"

	jwt "github.com/golang-jwt/jwt"
)

func JWTAuth(original func(w http.ResponseWriter, r *http.Request),
) func (w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return 
		}

		authHeaderParsed := strings.Split(authHeader[0], " ")
		if len(authHeaderParsed) != 2 || strings.ToLower(authHeaderParsed[0]) != "bearer" {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return 
		}

		if validateToken(authHeaderParsed[1]) {
			original(w,r)
		} else {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return 
		}
	}
}

func validateToken(accessToken string) bool {
	var myKey = []byte("letmeinkey")
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("could not validate auth token")
		}

		return myKey, nil 
	})
	if err != nil {
		return false
	}
	return token.Valid
}