package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/james-woo/wakeus/api/models"
	"github.com/james-woo/wakeus/api/utils"
	"net/http"
	"os"
	"strings"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			tokenHeader := r.Header.Get("Authorization")

			if tokenHeader == "" {
				respondWithError(w, false, "Missing auth token")
				return
			}

			split := strings.Split(tokenHeader, " ")
			if len(split) != 2 {
				respondWithError(w, false, "Invalid/Malformed auth token")
				return
			}

			tokenPart := split[1]
			tk := &models.Token{}

			token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("token_password")), nil
			})

			if err != nil {
				respondWithError(w, false, "Malformed auth token")
				return
			}

			if !token.Valid {
				respondWithError(w, false, "Token is not valid")
				return
			}

			next.ServeHTTP(w, r)
		},
	)
}

func respondWithError(w http.ResponseWriter, status bool, message string) {
	response := make(map[string]interface{})
	response = utils.Message(status, message)
	w.WriteHeader(http.StatusForbidden)
	w.Header().Add("Content-Type", "application/json")
	utils.Respond(w, response)
}