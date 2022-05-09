package routes

import (
	"crud-postgres/src/controllers"

	"github.com/gorilla/mux"
)

// Routes configure all routes
func Routes() *mux.Router {
	routes := mux.NewRouter()

	routes.HandleFunc("/", controllers.GetUsers).Methods("GET", "OPTIONS")
	routes.HandleFunc("/", controllers.CreateUser).Methods("POST", "OPTIONS")
	routes.HandleFunc("/", controllers.DeleteUser).Methods("DELETE", "OPTIONS")
	routes.HandleFunc("/", controllers.UpdateUser).Methods("PATCH", "OPTIONS")
	routes.HandleFunc("/login", controllers.Login).Methods("POST", "OPTIONS")
	routes.HandleFunc("/user", controllers.GetUser).Methods("GET", "OPTIONS")

	return routes
}
