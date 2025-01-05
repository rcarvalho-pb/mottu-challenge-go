package service

import (
	"github.com/rcarvalho-pb/mottu-broker_service/internal/config"
)

type BrokerService struct {
	UserSvc       *userService
	MotorcycleSvc *motorcycleService
	LocationSvc   *locationService
	AuthSvc       *authService
}

var addrs *config.Adresses

func New(addrPool *config.Adresses) *BrokerService {
	addrs = addrPool
	return &BrokerService{
		UserSvc:       newUserService(),
		MotorcycleSvc: newMotorcycleService(),
		LocationSvc:   newLocationService(),
		AuthSvc:       newAuthService(),
	}
}
