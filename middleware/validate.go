package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
)

func ValidateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// test_mode := req.Header.Get("Test")

		// if test_mode == "testing" {
		// 	next.ServeHTTP(w, req)
		// 	return
		// }

		authorizationHeader := req.Header.Get("Authorization")
		if authorizationHeader != "" {

			token, error := jwt.Parse(authorizationHeader, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return []byte("secret"), nil
			})
			if error != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(error.Error())
				return
			}
			if token.Valid {
				context.Set(req, "decoded", token.Claims)
				next.ServeHTTP(w, req)
				return
			} else {
				json.NewEncoder(w).Encode("Invalid authorization token")
				w.WriteHeader(498)
				return
			}

		} else {
			json.NewEncoder(w).Encode("An authorization header is required")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	})
}
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Max-Age", "86400")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, token,test")
		switch req.Method {
		case "OPTIONS":
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, req)
		return
	})
}
