package utils

import (
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

var Uid int

type MyCustomClaims struct {
	Uid int `json:"uid"`
	jwt.StandardClaims
}

// 生产token
func CreateToken(uid int, accessSecret string, accessExpire int64) (string, error) {
	claims := MyCustomClaims{Uid: uid}
	claims.StandardClaims.ExpiresAt = jwt.NewTime(float64(time.Now().Unix() + accessExpire))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(accessSecret))
	return ss, err
}

// 验证token
func CheckToken(tokenString string, accessSecret string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessSecret), nil
	})
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims.Uid, nil
	}
	return 0, err
}
