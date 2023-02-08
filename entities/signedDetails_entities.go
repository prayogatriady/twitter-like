package entities

import "github.com/dgrijalva/jwt-go"

type SignedDetails struct {
	User User
	jwt.StandardClaims
}
