package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID int `json:"user_id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"-"`
	Roles []Role `json:"roles"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy *int `json:"updated_by"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy *int `json:"deleted_by"`
}

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type UserClaims struct {
	UserID int `json:"user_id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Roles []Role `json:"roles"`
	jwt.RegisteredClaims
}