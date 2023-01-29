package models

import "github.com/golang-jwt/jwt"

type MyClaims struct {
	Uid                int64
	Username           string
	jwt.StandardClaims //StandardClaims
}
