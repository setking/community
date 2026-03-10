package models

import "github.com/dgrijalva/jwt-go"

// Define the CustomClaims struct for JWT claims
type CustomClaims struct {
	ID          uint64
	UserName    string
	AuthorityId uint
	jwt.StandardClaims
}
