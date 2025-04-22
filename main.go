package main

import (
	"log"
	"net/http"

	"github.com/MarcellinusAditya/go-jwt/controllers/authcontroller"
	"github.com/MarcellinusAditya/go-jwt/controllers/productcontroller"
	"github.com/MarcellinusAditya/go-jwt/middlewares"
	"github.com/MarcellinusAditya/go-jwt/models"
	"github.com/gorilla/mux"
)

func main() {
	models.ConnectDatabase()

	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/product", productcontroller.Index).Methods("GET")
	api.Use(middlewares.JWTMiddleware)
	
	log.Fatal(http.ListenAndServe(":8080",r))
	
}