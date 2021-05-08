package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// Claims custom token
type Claims struct {
	UserId   int64  `json:"userId"`   // 用户id
	Username string `json:"username"` // 用户名
	jwt.StandardClaims
}

// 生成token
func CreateToken(claims *Claims) (signedToken string, err error) {
	claims.ExpiresAt = time.Now().Add(time.Minute * 30).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte("secret"))
	return
}

// 验证token
func ValidateToken(signedToken string) (claims *Claims, success bool) {
	token, err := jwt.ParseWithClaims(signedToken, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected login method %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*Claims)
	if ok && token.Valid {
		success = true
		return
	}
	return
}
