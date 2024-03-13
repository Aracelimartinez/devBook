package router

import (
	"api/src/router/endpoints"
	"net/http"

	"github.com/gorilla/mux"
)

//Retorna um router com as rotas configuradas
func Generate() http.Handler {
 router := mux.NewRouter()
 return endpoints.Configure(router)
}
