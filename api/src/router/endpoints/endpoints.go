package endpoints

import (
	"api/src/middlewares"
	"net/http"

	"github.com/rs/cors"
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
func Configure(router *mux.Router) http.Handler {
	// Initialize CORS options
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	 handler := corsOptions.Handler(router)

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

	return handler
}
