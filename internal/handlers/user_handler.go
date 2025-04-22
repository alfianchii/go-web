package handlers

import (
	"errors"
	"go-web/internal/models"
	"go-web/internal/services"
	"go-web/internal/utils"
	"net/http"
)

var (
	ErrParseForm = errors.New("failed to parse form data")
	ErrGenerateJWT = errors.New("failed to generate JWT")
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
		utils.SendRes(res, ErrParseForm.Error(), http.StatusBadRequest, nil, err.Error())
		return
	}

	ipAddress := utils.GetClientIP(req)

	creds := models.LoginRequest{
		Username: req.FormValue("username"),
		Password: req.FormValue("password"),
	}

	token, err := h.userSvc.GenerateJWT(req.Context(), creds, ipAddress)
	if err != nil {
		utils.SendRes(res, ErrGenerateJWT.Error(), http.StatusUnauthorized, nil, err.Error())
		return
	}

	utils.SendRes(res, "Login successful", http.StatusOK, map[string]string{
		"token": token,
	}, "")
}