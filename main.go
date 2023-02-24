package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/karncomsci/golang-rest-api/db"
	"github.com/karncomsci/golang-rest-api/models"
	"github.com/karncomsci/golang-rest-api/routes"
)

func main() {
	db.DBConnection()

	//set model
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/GetUser/{id}", routes.GetUsersByIdHandler).Methods("GET")
	r.HandleFunc("/GetUsers", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/CreateUser", routes.CreateUserHandler).Methods("POST")
	r.HandleFunc("/DeleteUser/{id}", routes.DeleteUserHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
