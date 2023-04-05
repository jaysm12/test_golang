package middlewares

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type ErrorResponse struct {
	Success bool
	Msg     string
}

var excludeAuth = []string{"/api/user/register", "/api/user/login"}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for i, _ := range excludeAuth {
			if r.RequestURI == excludeAuth[i] {
				next.ServeHTTP(w, r)
				return
			}
		}
		reqToken := r.Header.Get("Authorization")
		if reqToken == "" {
			w.WriteHeader(http.StatusUnauthorized)

			response := ErrorResponse{
				Success: false,
				Msg:     "UNAUTHORIZED",
			}

			res, _ := json.Marshal(response)
			w.Write(res)
			return
		}

		tokenString := strings.Split(reqToken, "Bearer ")[1]

		token, _ := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)

				response := ErrorResponse{
					Success: false,
					Msg:     "UNAUTHORIZED",
				}

				res, _ := json.Marshal(response)
				w.Write(res)
			}

			return "", nil
		})

		if _, ok := token.Claims.(jwt.MapClaims); ok {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)

			response := ErrorResponse{
				Success: false,
				Msg:     "UNAUTHORIZED",
			}

			res, _ := json.Marshal(response)
			w.Write(res)
		}
	})
}
