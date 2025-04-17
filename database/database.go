package database

import (
	"github.com/dalobarahama/expense-tracker/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(sqlite.Open("expenses.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database!")
	}

	database.AutoMigrate(&models.Expense{})

	DB = database
}
