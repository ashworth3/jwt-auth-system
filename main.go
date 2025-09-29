package main

import (
	"log"
	"net/http"
	"jwt-auth-system/config"
	"jwt-auth-system/routes"
)

func main() {
	db := config.InitDB() // You define this
	r := routes.AuthRoutes(db)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
