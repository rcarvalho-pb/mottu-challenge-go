package controllers

import (
	"net/http"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/dtos"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/helper"
)

type authController struct{}

func newAuthController() *authController {
	return &authController{}
}

func (ac *authController) Authenticate(w http.ResponseWriter, r *http.Request) {
	var authRequest *dtos.AuthRequest
	if err := helper.ReadJson(w, r, &authRequest); err != nil {
		helper.ErrorJson(w, err, http.StatusBadRequest)
		return
	}
	tokenString, err := srv.AuthSvc.Authenticate(authRequest)
	if err != nil {
		helper.ErrorJson(w, err, http.StatusUnprocessableEntity)
		return
	}
	helper.WriteJson(w, http.StatusAccepted, tokenString)
}
