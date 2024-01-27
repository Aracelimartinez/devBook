package router

import (
	"api/src/router/endpoints"

	"github.com/gorilla/mux"
)

//Retorna um router com as rotas configuradas
func Generate() *mux.Router {
 router := mux.NewRouter()
 return endpoints.Configure(router)
}
