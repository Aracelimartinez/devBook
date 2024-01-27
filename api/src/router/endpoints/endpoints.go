package endpoints

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Representas todas as rotas da API
type Endopoints struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

// Coloca todas as rotas dentro do router
func Configure(router *mux.Router) *mux.Router {
	endpoints := userEndpoints
	endpoints = append(endpoints, loginEndpoint)
	endpoints = append(endpoints, postsEndpoints...)

	for _, endpoint := range endpoints {

		if endpoint.AuthRequired {
			router.HandleFunc(endpoint.URI, middlewares.Authenticate(endpoint.Function)).Methods(endpoint.Method)
		} else {
			router.HandleFunc(endpoint.URI, endpoint.Function).Methods(endpoint.Method)
		}
	}

	return router
}
