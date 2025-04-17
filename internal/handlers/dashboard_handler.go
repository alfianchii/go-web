package handlers

import (
	"errors"
	"go-web/internal/middlewares"
	"go-web/internal/models"
	"go-web/internal/services"
	"go-web/internal/utils"
	"net/http"
)

var (
	ErrFetchDashboardData = errors.New("")
)

type DashboardHdl interface {
	DashboardData(res http.ResponseWriter, req *http.Request)
}

type DashboardHdlImpl struct {
	dashboardSvc services.DashboardSvc
}

func NewDashboardHdl(dashboardSvc services.DashboardSvc) DashboardHdl {
	return &DashboardHdlImpl{
		dashboardSvc: dashboardSvc,
	}
}

func (h *DashboardHdlImpl) DashboardData(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	userClaims := ctx.Value(middlewares.UserClaimsKey).(*models.UserClaims)
	
	dashboardData, err := h.dashboardSvc.GetDashboardData(ctx, userClaims)
	if err != nil {
		utils.SendRes(res, ErrFetchDashboardData.Error(), http.StatusBadRequest, nil, err)
		return
	}

	utils.SendRes(res, "Dashboard data successfully fetched!", http.StatusOK, dashboardData, nil)
}