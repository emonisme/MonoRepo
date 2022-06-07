package main

import (
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"simpleapp/module/config"
)

func main() {
	restConfig := config.RestConfig{}
	err := envconfig.Process("rest", &restConfig)
	if err != nil {
		log.Fatal(err)
	}

	server, err := config.NewRestServer(restConfig)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":8080", server))
}
