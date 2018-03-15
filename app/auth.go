package app

import (
	"net/http"
	"fmt"
	"phonebook/controllers"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"os"
	"phonebook/models"
	"context"
)

var JwtMiddleWare = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		passThrough := []string{"/api/auth", "/api/account"}

		requestPath := r.URL.Path;
		for _, value := range passThrough {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string] interface{})
		headerString := r.Header.Get("Authorization")
		parts := strings.Split(headerString, " ")

		if strings.TrimSpace(headerString) == "" {
			response["status"] = false
			response["message"] = "Missing auth token"
			controllers.Respond(w, response)
			return
		}

		if len(parts) != 2 {
			response["status"] = false
			response["message"] = "Invalid token"
			controllers.Respond(w, response)
			return
		}

		tokenString := parts[1]

		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenString, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			response["status"] = false
			response["message"] = "Malformed token"
			response["err"] = err
			controllers.Respond(w, response)
			return
		}

		if !token.Valid {
			response["status"] = false
			response["message"] = "Token is invalid"
			controllers.Respond(w, response)
			return
		}

		fmt.Printf("%s => %d %s", "Decoded ==> ", tk.User, tk.Username)
		ctx := context.WithValue(r.Context(), "user", tk.User)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
