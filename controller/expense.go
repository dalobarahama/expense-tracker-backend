package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dalobarahama/expense-tracker/models"
)

var expenses []models.Expense

func GetExpenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expenses)
}

func CreateExpenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var expense models.Expense
	json.NewDecoder(r.Body).Decode(&expense)
	expense.ID = len(expenses) + 1
	expenses = append(expenses, expense)
	json.NewEncoder(w).Encode(expense)
}
