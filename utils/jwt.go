package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"new/models"
	"time"
)

var jwtKey = []byte("YSrDw8rTtKVH7Vuzf9GvC7eTCEesBYQs7sYkLyJvK8VwWbj3qZ\n")

type JwtCustClaims struct {
	ID   int64
	Name string
	jwt.RegisteredClaims
}

// GenerateToken 生成加密的 Token
func GenerateToken(user *models.UserLogin) (string, error) {
	iJwtCustClaims := JwtCustClaims{
		ID:   user.ID,
		Name: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iJwtCustClaims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", errors.New("生成 Token 错误")
	}
	return tokenStr, nil
}

// ParseToken 解析token
func ParseToken(tokenStr string) (JwtCustClaims, bool) {
	iJwtCustClaims := JwtCustClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &iJwtCustClaims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err == nil && !token.Valid {
		err = errors.New("invalid token")
	}
	return iJwtCustClaims, true
}
