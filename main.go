package main

import (
	"log"
	"net/http"

	"github.com/controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controllers.HomePage)
	router.HandleFunc("/user", controllers.CreateNewUsers).Methods("POST")
	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.GetNewUserID).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.UpdateUsers).Methods("PATCH")
	router.HandleFunc("/users/{id}", controllers.DeleteUsers).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
