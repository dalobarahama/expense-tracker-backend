package routes

import (
	"net/http"

	controllers "github.com/dalobarahama/expense-tracker/controller"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	}).Methods("GET")

	router.HandleFunc("/expenses", controllers.GetExpenses).Methods("GET")
	router.HandleFunc("/expenses", controllers.CreateExpenses).Methods("POST")

	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")

	return router
}
