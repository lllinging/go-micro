package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	//mux is the router for the application and will be returned to the main function to start the server with it
	mux := chi.NewRouter()

	// specify who is allowed to connect
	//mux.use is a middleware that will be executed before the request is passed to the handler which is the Broker function in this case
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},// Expose the headers				
		AllowCredentials: true,// Allow cookies
		MaxAge: 300,// Maximum value not ignored by any of major browsers
	}))

	//middleware.Heartbeat is a middleware that will be executed before the request is passed to the handler
	//check if the server is up and running
	mux.Use(middleware.Heartbeat("/ping"))

	mux.Post("/", app.Broker)

	mux.Post("/log-grpc", app.LogViaGRPC)

	mux.Post("/handle", app.HandleSubmission)


	return mux
}