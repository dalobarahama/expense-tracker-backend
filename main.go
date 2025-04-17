package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dalobarahama/expense-tracker/database"
	"github.com/dalobarahama/expense-tracker/routes"
)

func main() {
	database.Connect()

	router := routes.SetupRoutes()

	fmt.Println("✅ Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
