package rpc

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/rcarvalho-pb/mottu-motorcycle_service/internal/application/dtos"
	"github.com/rcarvalho-pb/mottu-motorcycle_service/internal/application/service"
)

type RPCServer struct {
	motorcycleService service.MotorcycleService
	Port              string
}

func New(service service.MotorcycleService, port string) *RPCServer {
	return &RPCServer{
		motorcycleService: service,
		Port:              port,
	}
}

func (r *RPCServer) RPCListen() error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", r.Port))
	if err != nil {
		return err
	}

	defer listen.Close()

	err = rpc.RegisterName("MotorcycleService", r)
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

func (r *RPCServer) GetMotorcycleById(motorcycleId *int64, entry *dtos.MotorcycleDTO) error {
	if motorcycleId == nil {
		return fmt.Errorf("motorcycle id can't be null")
	}
	motorcycle, err := r.motorcycleService.GetMotorcycleById(*motorcycleId)
	if err != nil {
		return err
	}

	*entry = *motorcycle

	return nil
}

func (r *RPCServer) CreateMotorcycle(request *dtos.NewMotorcycleRequest, _ *struct{}) error {
	if err := r.motorcycleService.CreateMotorcycle(request); err != nil {
		return err
	}

	return nil
}

func (r *RPCServer) UpdateMotorcycle(dto *dtos.MotorcycleDTO, _ *struct{}) error {
	if err := r.motorcycleService.UpdateMotorcycle(dto); err != nil {
		return err
	}

	return nil
}

func (r *RPCServer) DeleteMotorcycle(motorcycleId *int64, _ *struct{}) error {
	if motorcycleId == nil {
		return fmt.Errorf("motorcycle id can't be null")
	}

	if err := r.motorcycleService.DeleteMotorcycleById(*motorcycleId); err != nil {
		return err
	}

	return nil
}

func (r *RPCServer) GetAllActiveMotorcycles(_ struct{}, entry *[]*dtos.MotorcycleDTO) error {
	motorcycles, err := r.motorcycleService.GetAllActiveMotorcycles()
	if err != nil {
		return err
	}

	*entry = motorcycles

	return nil
}

func (r *RPCServer) GetAllMotorcycles(_ struct{}, entry *[]*dtos.MotorcycleDTO) error {
	motorcycles, err := r.motorcycleService.GetAllMotorcycles()
	if err != nil {
		return err
	}

	*entry = motorcycles

	return nil
}

func (r *RPCServer) GetAllMotorcyclesByYear(year *int64, entry *[]*dtos.MotorcycleDTO) error {
	if year == nil {
		return fmt.Errorf("year can't be null")
	}
	motorcycles, err := r.motorcycleService.GetMotorcyclesByYear(*year)
	if err != nil {
		return err
	}

	*entry = motorcycles

	return nil
}

func (r *RPCServer) GetAllMotorcyclesByModel(model *string, entry *[]*dtos.MotorcycleDTO) error {
	if model == nil {
		return fmt.Errorf("model can't be null")
	}

	motorcycles, err := r.motorcycleService.GetMotorcyclesByModel(*model)
	if err != nil {
		return err
	}

	*entry = motorcycles

	return nil
}

func (r *RPCServer) LocateMotorcycle(motorcycleId *int64, _ *struct{}) error {
	if motorcycleId == nil {
		return fmt.Errorf("motorcycle id can't be null")
	}

	if err := r.motorcycleService.LocateMotorcycle(*motorcycleId); err != nil {
		return err
	}

	return nil
}

func (r *RPCServer) UnlocateMotorcycle(motorcycleId *int64, _ *struct{}) error {
	if motorcycleId == nil {
		return fmt.Errorf("motorcycle id can't be null")
	}

	if err := r.motorcycleService.UnlocateMotorcycle(*motorcycleId); err != nil {
		return err
	}

	return nil
}
