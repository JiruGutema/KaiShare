package pkg

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jirugutema/gopastebin/internal/config"
)

var (
	jwtSecret     = []byte(config.LoadConfig().JWTSecret)
	refreshSecret = []byte(config.LoadConfig().RefreshSecret)
)

func GenerateJWT(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
		"iat":    time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenSignatureInvalid
		}

		return jwtSecret, nil
	})
}

func GenerateRefreshToken(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 days
		"iat":    time.Now().Unix(),
		"type":   "refresh",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(refreshSecret)
}

func ValidateRefreshToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenSignatureInvalid
		}
		return refreshSecret, nil
	})
}

func GetClaims(token *jwt.Token) jwt.MapClaims {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims
	}
	return nil
}

func GetUserIDFromToken(token *jwt.Token) (uuid.UUID, error) {
	claims := GetClaims(token)
	if claims == nil {
		return uuid.UUID{}, fmt.Errorf("invalid claims")
	}

	userIDString, ok := claims["userID"].(string)
	if !ok {
		return uuid.UUID{}, fmt.Errorf("userID missing or not a string")
	}

	return uuid.Parse(userIDString)
}
