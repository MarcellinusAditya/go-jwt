package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("wiojfioq9uq39u1903iocn28h")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}