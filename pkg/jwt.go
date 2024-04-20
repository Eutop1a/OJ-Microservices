package pkg

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	Username  string `json:"username"`
	TokenType string `json:"token-type"`
	jwt.StandardClaims
}
type Token struct {
	TokenString string // token
}

type Result struct {
	UserName string // 返回的username
	Status   string // 状态：
	// case 0: "" (success)
	// case 1: "Token has expired"
	// case 2: "Error parsing token:"
}

const TokenTime = time.Hour * 24

var Secret = []byte("Author:Eutop1a")

const Author = "Eutop1a"

// GenerateToken 生成Token
func GenerateToken(username string) string {
	c := MyClaims{
		username, "AccessToken", jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenTime).Unix(),
			Issuer:    Author,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	res, _ := token.SignedString(Secret)
	return res
}

func JudgeAccessToken(tokenString string) (username string, ok bool) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return "", false // Token has expired
			}
		}
		return "", false // Error parsing token
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims.Username, true
	} else {
		return "", false // Error parsing token
	}
}

func ParseToken(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})

	if err != nil {
		return err
	}

	return nil
}
