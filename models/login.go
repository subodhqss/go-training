package models

import "github.com/dgrijalva/jwt-go"

type Login struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Token     string `json:"token"`
	ExpiresAt string `json:"expires_at"`
	
	jwt.StandardClaims
}
