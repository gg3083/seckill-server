package util

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

var jwtSecret []byte

type Claims struct {
	PkId     interface{} `json:"pk_id"`
	UserName string      `json:"user_name"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(pkId interface{}, userName string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(60) * time.Minute) // 1一个小时的过期时间

	claims := Claims{
		pkId,
		userName,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "quick-pass",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	//err = token_cache.SetCacheToken(agency, username, token)
	if err != nil {
		log.Println("token_cache.SetCacheToken:", err)
	}
	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
