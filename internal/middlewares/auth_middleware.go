package middlewares

import (
	"context"
	"errors"
	"go-web/configs"
	"go-web/internal/repositories"
	"go-web/internal/services"
	"go-web/internal/utils"
	"net/http"
)

type contextKey string
const UserClaimsKey contextKey = "userClaims"

var (
	ErrTokenFormat = errors.New("invalid token format")
	ErrBlacklistedToken = errors.New("unauthorized: Token is blacklisted")
	ErrGetBlacklistedToken= errors.New("failed to get blacklisted token")
	ErrToken = errors.New("unauthorized: Invalid token")
	ErrInsufficientPerms = errors.New("forbidden: Insufficient permissions")
)

func AuthMiddleware(requiredRole string, userSvc services.UserSvc, sessionRepo repositories.SessionRepo) func (http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			tokenString, err := utils.GetBearerToken(req.Header.Get("Authorization"))
			if err != nil {
				utils.SendRes(res, ErrTokenFormat.Error(), http.StatusUnauthorized, nil, err.Error())
				return
			}

			isBlacklisted, err := sessionRepo.IsTokenBlacklisted(req.Context(), tokenString)
			if err != nil {
				utils.SendRes(res, ErrGetBlacklistedToken.Error(), http.StatusUnauthorized, nil, err.Error())
				return
			}
			if isBlacklisted {
				utils.SendRes(res, ErrBlacklistedToken.Error(), http.StatusUnauthorized, nil, "")
				return
			}

			userClaims, err := utils.ValidateJWT(tokenString, configs.GetENV("JWT_SECRET"))
			if err != nil {
				utils.SendRes(res, ErrToken.Error(), http.StatusUnauthorized, nil, err.Error())
				return
			}

			hasRole := false
			for _, role := range userClaims.Roles {
				if role.Name == requiredRole {
					hasRole = true
					break
				}
			}

			if !hasRole {
				utils.SendRes(res, ErrInsufficientPerms.Error(), http.StatusForbidden, nil, "")
			}

			ctx := context.WithValue(req.Context(), UserClaimsKey, userClaims)
			next.ServeHTTP(res, req.WithContext(ctx))
		})
	}
}