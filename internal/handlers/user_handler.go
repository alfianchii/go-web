package handlers

import (
	"errors"
	"go-web/internal/models"
	"go-web/internal/services"
	"go-web/internal/utils"
	"net/http"
)

var (
	ErrFailedParseForm = errors.New("failed to parse form data")
	ErrFailedGenerateJWT = errors.New("failed to generate JWT")
)

type UserHdl interface {
	Login(res http.ResponseWriter, req *http.Request)
}

type UserHdlImpl struct {
	userSvc services.UserSvc
}

func NewUserHdl(userSvc services.UserSvc) UserHdl {
	return &UserHdlImpl{
		userSvc: userSvc,
	}
}

func (h *UserHdlImpl) Login(res http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		utils.SendRes(res, ErrFailedParseForm.Error(), http.StatusBadRequest, nil, err)
		return
	}

	ipAddress := utils.GetClientIP(req)

	creds := models.LoginRequest{
		Username: req.FormValue("username"),
		Password: req.FormValue("password"),
	}

	token, err := h.userSvc.GenerateJWT(req.Context(), creds, ipAddress)
	if err != nil {
		utils.SendRes(res, ErrFailedGenerateJWT.Error(), http.StatusUnauthorized, nil, err)
		return
	}

	utils.SendRes(res, "Login successful", http.StatusOK, map[string]string{
		"token": token,
	}, nil)
}