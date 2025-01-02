package service

import "github.com/rcarvalho-pb/mottu-broker_service/internal/config"

type BrokerService struct {
	userSvc       *UserService
	motorcycleSvc *MotorcycleService
	locationSvc   *LocationService
}

var addrs *config.Adresses

func New(addrs *config.Adresses) *BrokerService {
	addrs = addrs
	return &BrokerService{}
}
