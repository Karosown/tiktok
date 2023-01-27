package models

import "github.com/golang-jwt/jwt"

type Jwt struct {
	Uid      int64
	Username string
	jwt.StandardClaims
}
