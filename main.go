package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dalobarahama/expense-tracker/routes"
)

func main() {
	router := routes.SetupRoutes()

	fmt.Println("✅ Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
