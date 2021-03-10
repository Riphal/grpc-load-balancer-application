package jwt

import "github.com/dgrijalva/jwt-go"

// Claims are custom claims extending default ones.
type Claims struct {
	ID  				string `json:"id"`
	Email  				string `json:"email"`
	jwt.StandardClaims
}
