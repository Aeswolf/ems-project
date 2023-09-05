package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/Aeswolf/equipment-database-management/api"
	"github.com/Aeswolf/equipment-database-management/routes"
)

// function to start up a server
func Run(a *api.APIServer) error {
	r := mux.NewRouter()

	routes.Access(a, r)

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut},
		AllowedHeaders: []string{"Content-Type"},
	})

	corsHandler := cors.Handler(r)

	log.Printf("Server started running on port : %s\n", a.Port)
	return http.ListenAndServe(":"+a.Port, corsHandler)
}
