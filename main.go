package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vernandodev/go-restapi-jwt-mux/controllers/authcontrollers"
	"github.com/vernandodev/go-restapi-jwt-mux/models"
)

func main() {
	models.ConnectDatabase()
	route := mux.NewRouter()

	route.HandleFunc("/login", authcontrollers.Login).Methods("POST")
	route.HandleFunc("/register", authcontrollers.Register).Methods("POST")
	route.HandleFunc("/logout", authcontrollers.Logout).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", route))
}
