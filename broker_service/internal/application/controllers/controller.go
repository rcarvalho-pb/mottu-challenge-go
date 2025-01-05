package controllers

import "github.com/rcarvalho-pb/mottu-broker_service/internal/application/service"

var srv *service.BrokerService

func StartController(serv *service.BrokerService) {
	srv = serv
}
