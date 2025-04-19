package routes

import (
	controllers "github.com/dalobarahama/expense-tracker/controller"
	"github.com/dalobarahama/expense-tracker/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	authorized := router.Group("/")
	authorized.Use(middlewares.AuthMiddleware())
	{
		authorized.GET("/expenses", controllers.GetExpenses)
		authorized.POST("/expenses", controllers.CreateExpenses)
	}

	return router
}
