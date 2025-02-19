package middleware

import (
	"errors"
	"Go-Tutorial-Api/api"
	"Go-Tutorial-Api/internal/tools"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var UnauthorizedError = errors.New("Invalid token/user credentials")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token string = r.Header.Get("Authorization")
		var errorMsg error
		if username == "" || token == "" {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}
		var database *tools.DatabaseInterface
		database, errorMsg = tools.NewDatabase()
		if errorMsg != nil {
			api.InternalErrorHandler(w)
			return
		}
		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if (loginDetails == nil || token != (*loginDetails).AuthToken) {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}
		next.ServeHTTP(w, r)
	})
}
