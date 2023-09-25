package utils

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"time"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	Id        uint   `json:"id"`
	Username  string `json:"username"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// GenerateToken 颁发token
func GenerateToken(id uint, username string, authority int) Token {
	
	var tokenObj Token
	for i := 0; i < 2; i++ {
		nowTime := time.Now()
		var expireTime time.Time
		if i == 0 {
			expireTime = nowTime.Add(80 * time.Minute)
		} else {
			expireTime = nowTime.Add(24 * time.Hour)
		}
		
		claims := Claims{
			Id:        id,
			Username:  username,
			Authority: authority,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expireTime.Unix(),
				Issuer:    "meeting_booking",
			},
		}
		
		tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token, err := tokenClaims.SignedString(jwtSecret)
		if err != nil {
			log.Panicln("Error signing")
		}
		if i == 0 {
			tokenObj.AccessToken = token
		} else {
			tokenObj.RefreshToken = token
		}
	}
	return tokenObj
}

//ParseToken  验证用户token
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
