package framework

import (
	"net/http"
)

type FunctionsAuth interface {
	AuthenticationMiddleware(apiUser string, apiPassword string, next http.Handler) http.Handler
}

type FunctionsAuthentication struct{}

func (f *FunctionsAuthentication) AuthenticationMiddleware(apiUser string, apiPassword string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqUser, reqPassword, isOk := r.BasicAuth()
		if isOk && reqUser == apiUser && reqPassword == apiPassword {
			next.ServeHTTP(w, r)
			return
		}
		errorMessage := "Acceso no autorizado"
		http.Error(w, errorMessage, http.StatusUnauthorized)
		return
	})
}
