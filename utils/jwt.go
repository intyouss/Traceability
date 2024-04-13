package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type JwtClaims struct {
	ID     uint
	Name   string
	RoleID uint
	jwt.RegisteredClaims
}

// GenerateToken 生成token
func GenerateToken(id uint, name string, roleId uint) (string, error) {
	var secretKey = []byte(viper.GetString("jwt.secretKey"))
	jwtClaim := JwtClaims{
		ID:     id,
		Name:   name,
		RoleID: roleId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.tokenExpire") * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim).SignedString(secretKey)
}

// ParseToken 解析token
func ParseToken(tokenStr string) (*JwtClaims, error) {
	var secretKey = []byte(viper.GetString("jwt.secretKey"))
	jwtClaim := &JwtClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, jwtClaim, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return jwtClaim, nil
}
