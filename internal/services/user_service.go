package services

import (
	"context"
	"errors"
	"go-web/configs"
	"go-web/internal/models"
	"go-web/internal/repositories"
	"go-web/internal/utils"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidUserPass = errors.New("invalid username or password")
	ErrFailedUUIDForSession = errors.New("failed to generate UUID for session")
)

type UserSvc interface {
	GenerateJWT(ctx context.Context, creds models.LoginRequest, ipAddress string) (string, error)
}

type UserSvcImpl struct {
	userRepo repositories.UserRepo
	sessionRepo repositories.SessionRepo
}

func NewUserSvc(userRepo repositories.UserRepo, sessionRepo repositories.SessionRepo) UserSvc {
	return &UserSvcImpl{
		userRepo: userRepo,
		sessionRepo: sessionRepo,
	}
}

func (s *UserSvcImpl) GenerateJWT(ctx context.Context, creds models.LoginRequest, ipAddress string) (string, error) {
	user, err := s.userRepo.FindByUsernameWithRoles(ctx, creds.Username)
	if err != nil {
		return "", errors.New(err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		return "", ErrInvalidUserPass
	}

	jwt, err := utils.GenerateJWT(user, configs.GetENV("JWT_SECRET"))
	if err != nil {
		return "", err
	}

	rowID, err := uuid.NewRandom()
	if err != nil {
		return "", ErrFailedUUIDForSession
	}

	session := models.Session{
		RowID: rowID,
		UserID: user.ID,
		Token: jwt.Token,
		ExpiresAt: jwt.ExpiresAt.Time,
		CreatedAt: time.Now().Local(),
		IPAddress: ipAddress,
		IsBlacklisted: false,
	}

	s.sessionRepo.StoreSession(ctx, session)

	return jwt.Token, nil
}