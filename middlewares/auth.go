package middlewares

import (
	"net/http"
	"time"

	controllers "github.com/dalobarahama/expense-tracker/controller"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret_key")

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		claims := &controllers.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid || claims.ExpiresAt.Time.Before(time.Now()) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"token": "Invalid or expired access token"})
			return
		}

		ctx.Set("username", claims.Username)
		ctx.Set("user_id", claims.UserID)
		ctx.Next()
	}
}
