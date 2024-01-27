package middlewares

import (
	"api/src/authentication"
	"api/src/responses"
	"net/http"
)

// Verifica se o usuário fazendo a requisição está autenticado
func Authenticate(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.ValidateToken(r); err != nil {
			responses.Err(w, http.StatusUnauthorized, err)
			return
		}
		nextFunction(w, r)
	}
}
