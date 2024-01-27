package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := router.Generate()
	fmt.Printf("Escutando na porta %s\n", config.GlobalConfig.APIPort)
	
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.GlobalConfig.APIPort), router))
}
