package utils

import (
	"errors"
	"go-web/configs"
	"go-web/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrFailedToSignToken = errors.New("failed to sign token")
	ErrEmptyAuthHeader = errors.New("unauthorized; authorization header is empty")
	ErrInvalidBearerToken = errors.New("unathorized; authorization header is not a Bearer token")
	ErrTokenInvalid = errors.New("invalid token")
	ErrTokenInvalidClaims = errors.New("invalid token claims")
)

const bearerPrefix = "Bearer "
const bearerPrefixLength = len(bearerPrefix)

type GeneratedJWT struct {
	Token string `json:"token"`
	ExpiresAt *jwt.NumericDate `json:"expires_at"`
}

func GenerateJWT(user *models.User, secret string) (GeneratedJWT, error) {
	tokenExp := jwt.NewNumericDate(time.Now().Add(configs.TokenDuration))

	claims := models.UserClaims{
		UserID: user.ID,
		Name: user.Name,
		Username: user.Username,
		Roles: user.Roles,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: user.Name,
			Subject: user.Username,
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: tokenExp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return GeneratedJWT{}, ErrFailedToSignToken
	}

	return GeneratedJWT{
		Token: signedToken,
		ExpiresAt: tokenExp,
	}, nil
}

func ValidateJWT(tokenString string, secret string) (*models.UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, ErrTokenInvalid
	}

	claims, ok := token.Claims.(*models.UserClaims)
	if !ok {
		return nil, ErrTokenInvalidClaims
	}

	return claims, nil
}

func GetBearerToken(authHeader string) (string, error) {
	if authHeader == "" {
		return "", ErrEmptyAuthHeader
	}

	if len(authHeader) < bearerPrefixLength || authHeader[:bearerPrefixLength] != bearerPrefix {
		return "", ErrInvalidBearerToken
	}

	return authHeader[bearerPrefixLength:], nil
}