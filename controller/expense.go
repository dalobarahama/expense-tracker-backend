package controllers

import (
	"net/http"

	"github.com/dalobarahama/expense-tracker/database"
	"github.com/dalobarahama/expense-tracker/models"
	"github.com/gin-gonic/gin"
)

func GetExpenses(ctx *gin.Context) {
	var expenses []models.Expense

	result := database.DB.Find(&expenses)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, expenses)
}

func CreateExpenses(ctx *gin.Context) {
	var expense models.Expense

	if err := ctx.ShouldBindJSON(&expense); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&expense)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, expense)
}
