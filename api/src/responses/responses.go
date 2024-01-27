package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// Retorna uma resposta em JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatalln(err)
		}
	}
}

// Retorna um erro em formato JSON
func Err(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct{
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}
