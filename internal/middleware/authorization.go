package middleware

import (
	"errors"
	"net/http"
	"github.com/ShinichiShi/lessGO/api"
	"github.com/ShinichiShi/lessGO/internal/tools"
	log "github.com/sirupsen/logrus"
)
var UnAuthorizedError = errors.New("Invalid Username or token")

func Authorization(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error
		if username=="" || token =="" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w,UnAuthorizedError)

			return
		}
		var database *tools.DatabaseInterface
		database,err = tools.NewDatabase()
		if err!=nil{
			api.InternalErrorHandler(w)
			return
		}
		var loginDetails *tools.loginDetails
		loginDetails = (*database).GetUSerLoginDetails(username)
		if(loginDetails==nil || (token != (*&loginDetails).AuthToken)){
			log.Error((UnAuthorizedError))
			api.RequestErrorHandler(w,UnAuthorizedError) 
			return
		}
		next.ServeHTTP(w,r)
	})
}