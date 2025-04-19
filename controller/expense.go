package controllers

import (
	"fmt"
	"net/http"

	"github.com/dalobarahama/expense-tracker/database"
	"github.com/dalobarahama/expense-tracker/models"
	"github.com/gin-gonic/gin"
)

func GetExpenses(ctx *gin.Context) {
	userId, _ := ctx.Get("user_id")

	var expenses []models.Expense
	if err := database.DB.Where("user_id = ?", userId.(uint)).Find(&expenses).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get expenses"})
		return
	}

	fmt.Println("userId ", userId)

	ctx.JSON(http.StatusOK, expenses)
}

func CreateExpenses(ctx *gin.Context) {
	userId, _ := ctx.Get("user_id")

	var expense models.Expense
	if err := ctx.BindJSON(&expense); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expense.UserID = userId.(uint)

	if err := database.DB.Create(&expense).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create expenses"})
		return
	}

	fmt.Println("userId ", userId)

	ctx.JSON(http.StatusCreated, expense)
}
