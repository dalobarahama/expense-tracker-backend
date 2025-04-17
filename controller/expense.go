package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dalobarahama/expense-tracker/database"
	"github.com/dalobarahama/expense-tracker/models"
)

func GetExpenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var expenses []models.Expense
	database.DB.Find(&expenses)
	json.NewEncoder(w).Encode(expenses)
}

func CreateExpenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var expense models.Expense
	json.NewDecoder(r.Body).Decode(&expense)
	database.DB.Create(&expense)
	json.NewEncoder(w).Encode(expense)
}
