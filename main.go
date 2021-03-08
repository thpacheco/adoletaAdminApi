package main

import (
	authController "adoletaAdminApi/controllers/auth"
	usersController "adoletaAdminApi/controllers/users"
	jwtService "adoletaAdminApi/jwtSecurity"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	s := route.PathPrefix("/api").Subrouter() //Base Path
	//Routes

	// s.Handle("/createProfile", jwtService.IsAuthorized(userController.CreateProfile)).Methods("POST")
	s.Handle("/getAllUsers", jwtService.IsAuthorized(usersController.GetAllUsers)).Methods("GET")
	// s.Handle("/getUserProfile", jwtService.IsAuthorized(userController.GetUserProfile)).Methods("POST")
	// s.Handle("/updateProfile", jwtService.IsAuthorized(userController.UpdateProfile)).Methods("PUT")
	// s.Handle("/deleteProfile/{id}", jwtService.IsAuthorized(userController.DeleteProfile)).Methods("DELETE")
	s.HandleFunc("/createProfile", usersController.CreateProfile).Methods("POST")
	s.HandleFunc("/auth", authController.Auth).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", s)) // Run Server

}
