package jwt

import (
	"fmt"
	"strings"
	"time"

	"github.com/RivGames/my-knowledge-base/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func CreateToken(id uint, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	secret := config.Envs.JWTSecretKey
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
}

func GetUserID(c echo.Context) (uint, error) {
	token := strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")
	jwtToken, err := VerifyToken(token)
	if err != nil {
		return 0, err
	}

	return uint(jwtToken.Claims.(jwt.MapClaims)["id"].(float64)), nil
}
