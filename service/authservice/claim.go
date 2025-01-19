package authservice

import (
	"github.com/golang-jwt/jwt/v5"
	"main.go/entity"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID uint        `json:"user_id"`
	Role   entity.Role `json:"role"`
}
