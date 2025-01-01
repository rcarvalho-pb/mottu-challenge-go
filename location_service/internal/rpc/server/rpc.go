package rpc_server

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/rcarvalho-pb/mottu-location_service/internal/application/dtos"
	"github.com/rcarvalho-pb/mottu-location_service/internal/application/service"
)

type RPCServer struct {
	locationService *service.LocationService
	Port            string
}

func New(locationService *service.LocationService, port string) *RPCServer {
	return &RPCServer{
		locationService: locationService,
		Port:            port,
	}
}

func (r *RPCServer) RPCListen() error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", r.Port))
	if err != nil {
		return err
	}

	defer listen.Close()

	err = rpc.RegisterName("LocationService", r)
	if err != nil {
		fmt.Println(err)
	}

	for {
		rpcConn, err := listen.Accept()
		if err != nil {
			return err
		}

		go rpc.ServeConn(rpcConn)
	}
}

func (r *RPCServer) GetLocationById(locationId *int64, entry *dtos.LocationDTO) error {
	if locationId == nil {
		return fmt.Errorf("location id can't be null")
	}

	location, err := r.locationService.GetLocationById(*locationId)
	if err != nil {
		return err
	}

	*entry = *location

	return nil
}

func (r *RPCServer) CreateLocatioin(location *dtos.NewLocationDTO, _ *struct{}) error {
	if location == nil {
		return fmt.Errorf("location can't be null")
	}

	if err := r.locationService.CreateLocation(location); err != nil {
		return err
	}

	return nil
}

func (r *RPCServer) UpdateLocation(location *dtos.LocationDTO, _ *struct{}) error {
	if location == nil {
		return fmt.Errorf("location can't be null")
	}

	if err := r.locationService.UpdateLocation(location); err != nil {
		return err
	}

	return nil
}

func (r *RPCServer) EndLocation(locationId *int64, _ *struct{}) error {
	if locationId == nil {
		return fmt.Errorf("location id can't be null")
	}

	if err := r.locationService.EndLocation(*locationId); err != nil {
		return err
	}

	return nil
}

func (r *RPCServer) GetAllLocations(_ struct{}, entry *[]*dtos.LocationDTO) error {
	locations, err := r.locationService.GetAllLocations()
	if err != nil {
		return err
	}

	*entry = locations

	return nil
}

func (r *RPCServer) GetAllActiveLocations(_ *struct{}, entry *[]*dtos.LocationDTO) error {
	locations, err := r.locationService.GetAllActiveLocations()
	if err != nil {
		return err
	}

	*entry = locations

	return nil
}

func (r *RPCServer) GetLocationsByUserId(userId *int64, entry *[]*dtos.LocationDTO) error {
	if userId == nil {
		return fmt.Errorf("user id can't be null")
	}
	locations, err := r.locationService.GetLocationByUserId(*userId)
	if err != nil {
		return err
	}

	*entry = locations

	return nil
}

func (r *RPCServer) GetLocationsByMotorcycleId(motorcycleId *int64, entry *[]*dtos.LocationDTO) error {
	if motorcycleId == nil {
		return fmt.Errorf("motorcycle id can't be null")
	}
	locations, err := r.locationService.GetLocationByMotorcycleId(*motorcycleId)
	if err != nil {
		return err
	}

	*entry = locations

	return nil
}
