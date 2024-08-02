package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {}

func main() {
	app := Config{}

	log.Printf("Starting broker service on port %s\n", webPort)

	// define http server
	//&http.Server is a struct that contains the configuration of the server
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),// Address to listen on
		Handler: app.routes(),// Handler to invoke, http.DefaultServeMux if nil
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}