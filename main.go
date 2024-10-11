package main

import (
	"invite-wed/controllers/authcontroller"
	"invite-wed/middlewares"
	"invite-wed/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	models.ConnectDB()

	r := mux.NewRouter()

	r.HandleFunc("/api/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/api/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/api/forgot-password/send-email", authcontroller.ForgotPasswordEmail).Methods("POST")
	r.HandleFunc("/api/forgot-password/reset", authcontroller.ForgotPasswordReset).Methods("POST")

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middlewares.JWTMiddleware)
	api.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
