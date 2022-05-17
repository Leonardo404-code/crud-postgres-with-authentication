package main

import (
	"crud-postgres/src/config"
	"crud-postgres/src/database"
	"crud-postgres/src/routes"
	"log"
	"net/http"
)

func init() {
	config.LoadDotEnv()

	database.Connect()
}

func main() {
	routes := routes.Routes()

	log.Println("Server Started at http://localhost:3000")

	if err := http.ListenAndServe(":3000", routes); err.Error() != "" {
		log.Fatalf("error %s", err)
	}
}
