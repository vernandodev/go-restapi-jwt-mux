package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vernandodev/go-restapi-jwt-mux/controllers/authcontrollers"
	"github.com/vernandodev/go-restapi-jwt-mux/controllers/productcontrollers"
	"github.com/vernandodev/go-restapi-jwt-mux/middlewares"
	"github.com/vernandodev/go-restapi-jwt-mux/models"
)

func main() {
	models.ConnectDatabase()
	route := mux.NewRouter()

	// ROUTE AUTH
	route.HandleFunc("/login", authcontrollers.Login).Methods("POST")
	route.HandleFunc("/register", authcontrollers.Register).Methods("POST")
	route.HandleFunc("/logout", authcontrollers.Logout).Methods("GET")

	// sub router / meng grouping routes
	api := route.PathPrefix("/api").Subrouter()

	// ROUTE PRODUCT
	api.HandleFunc("/products", productcontrollers.Index).Methods("GET")

	// untuk menggunakan middleware
	api.Use(middlewares.ProductMiddleware)

	log.Fatal(http.ListenAndServe(":8080", route))
}
