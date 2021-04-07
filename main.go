package main

import (
	authController "adoletaAdminApi/controllers/auth"
	usersController "adoletaAdminApi/controllers/users"
	jwtService "adoletaAdminApi/jwtSecurity"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	route := mux.NewRouter()

	s := route.PathPrefix("/api").Subrouter() //Base Path

	//Routes
	//s.Handle("/createProfile", jwtService.IsAuthorized(usersController.CreateProfile)).Methods("POST")
	//s.Handle("/getAllUsers", jwtService.IsAuthorized(usersController.GetAllUsers)).Methods("GET")
	s.Handle("/getUserProfile", jwtService.IsAuthorized(usersController.GetUserProfile)).Methods("POST")
	s.Handle("/updateProfile", jwtService.IsAuthorized(usersController.UpdateProfile)).Methods("PUT")
	s.Handle("/deleteProfile/{id}", jwtService.IsAuthorized(usersController.DeleteProfile)).Methods("DELETE")
	s.HandleFunc("/auth", authController.Auth).Methods("POST")

	s.HandleFunc("/getAllUsers", usersController.GetAllUsers).Methods("GET")
	s.HandleFunc("/createProfile", usersController.CreateProfile).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"Accept", "content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})

	handler := c.Handler(route)

	log.Fatal(http.ListenAndServe(GetPort(), handler)) // Run Server

}

func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "8000"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
