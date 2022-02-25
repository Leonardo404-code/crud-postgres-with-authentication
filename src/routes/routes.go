package routes

import (
	"crud-postgres/src/controllers"

	"github.com/gorilla/mux"
)

// Configure all routes
func Routes() *mux.Router {
	routes := mux.NewRouter()

	routes.HandleFunc("/", controllers.GetUsers).Methods("GET", "OPTIONS")
	routes.HandleFunc("/", controllers.CreateUser).Methods("POST", "OPTIONS")
	routes.HandleFunc("/{id}", controllers.DeleteUser).Methods("DELETE", "OPTIONS")
	routes.HandleFunc("/{id}", controllers.UpdateUser).Methods("PATCH", "OPTIONS")

	return routes
}
