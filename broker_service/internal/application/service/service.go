package service

import (
	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/dtos"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/config"
)

type BrokerService struct {
	userSvc       *userService
	motorcycleSvc *motorcycleService
	locationSvc   *locationService
	authSvc       *authService
}

var addrs *config.Adresses

func New(addrPool *config.Adresses) *BrokerService {
	addrs = addrPool
	return &BrokerService{
		userSvc:       newUserService(),
		motorcycleSvc: newMotorcycleService(),
		locationSvc:   newLocationService(),
		authSvc:       newAuthService(),
	}
}

func (bs *BrokerService) Authenticate(authRequest *dtos.AuthRequest) (*string, error){
	user, err := bs.userSvc.getUserByUsername(authRequest.Username)
	if err != nil {
		return nil, err
	}
	bs.userSvc.
}
