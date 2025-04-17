package services

import (
	"context"
	"go-web/internal/models"
)

type DashboardSvc interface {
	GetDashboardData(ctx context.Context, userClaims *models.UserClaims) (*models.UserClaims, error)
}

type DashboardSvcImpl struct {}

func NewDashboardSvc() DashboardSvc {
	return &DashboardSvcImpl{}
}

func (s *DashboardSvcImpl) GetDashboardData(ctx context.Context, userClaims *models.UserClaims) (*models.UserClaims, error) {
	return userClaims, nil
}