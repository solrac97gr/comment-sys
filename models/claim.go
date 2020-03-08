package models

import jwt "github.com/dgrijalva/jwt-go"

// Claim Token for user
type Claim struct {
	User `json:"user"`
	jwt.StandardClaims
}
