package main

import (
	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/service"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/config"
)

func main() {
	config.Start()
	brokerService := service.New(config.Addresses)
}
