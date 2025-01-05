package controllers

import "github.com/rcarvalho-pb/mottu-broker_service/internal/application/service"

var srv *service.BrokerService

type Controller struct {
	AuthController *authController
	UserController *userController
}

func NewController(serv *service.BrokerService) *Controller {
	srv = serv
	return &Controller{
		AuthController: newAuthController(),
		UserController: newUserController(),
	}
}
