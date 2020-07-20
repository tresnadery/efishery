package domain

import (
	"time"
	jwt "github.com/dgrijalva/jwt-go"	
)

type JwtToken struct {
	PhoneNumber string `json:"phone_number"`
	Name  string    `json:"name"`
	RoleName string `json:"role_name"`
	CreatedAt time.Time `json:"created_at"`
	*jwt.StandardClaims
}