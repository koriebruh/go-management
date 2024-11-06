package cnf

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("fIDJA90wju190djkqqlwpwqqwieuhd90q32")

type JWTClaim struct {
	UserID   int `json:"id"`
	UserName string
	jwt.RegisteredClaims
}
