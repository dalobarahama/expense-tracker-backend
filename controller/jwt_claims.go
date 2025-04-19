package controllers

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Username string `json:"username"`
	UserID   uint   `json:"user_id"`
	jwt.RegisteredClaims
}
