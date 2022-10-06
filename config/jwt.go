package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("js9dr98f8834u38f383h8y33ds")

// struct untuk menyimpan username, informasi expired, yg mengeluarkan token

type JWTclaim struct {
	Username string
	jwt.RegisteredClaims
}
