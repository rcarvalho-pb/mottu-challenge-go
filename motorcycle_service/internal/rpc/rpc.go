package rpc

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/rcarvalho-pb/mottu-motorcycle_service/internal/application/services"
)

type RPCServer struct {
	motorcycleService services.MotorcycleService
	Port              string
}

func New(service services.MotorcycleService, port string) *RPCServer {
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

func (r *RPCServer) GetUserById(userId int64, reply *dtos.UserDTO) error {
	user, err := r.motorcycleService.GetUserById(userId)
	if err != nil {
		return fmt.Errorf("error getting user by id: %s\n", err)
	}

	*reply = *user

	return err
}

func (r *RPCServer) GetMotorcycleByUsername(username string, reply *dtos.UserDTO) error {
	user, err := r.motorcycleService.GetUserByUsername(username)
	if err != nil {
		return fmt.Errorf("error getting user by username: %s\n", err)
	}

	*reply = *user
	return err
}

func (r *RPCServer) GetAllActiveUsers(_ struct{}, reply *[]*dtos.UserDTO) error {
	users, err := r.motorcycleService.GetAllActiveUsers()
	if err != nil {
		return fmt.Errorf("error getting all active users: %s\n", err)
	}

	*reply = users
	return err
}

func (r *RPCServer) GetAllUsers(_ struct{}, reply *[]*dtos.UserDTO) error {
	users, err := r.motorcycleService.GetAllUsers()
	if err != nil {
		return fmt.Errorf("error getting all users: %s\n", err)
	}

	*reply = users
	return err
}

func (r *RPCServer) DeactivateUser(userId int64, reply *bool) error {
	if err := r.motorcycleService.DeactivateUserById(userId); err != nil {
		return fmt.Errorf("error deactivating user [%d]: %s\n", userId, err)
	}

	*reply = true

	return nil
}

func (r *RPCServer) ReactivateUser(userId int64, reply *bool) error {
	if err := r.motorcycleService.ReactivateUserById(userId); err != nil {
		return fmt.Errorf("error reactivating user [%d]: %s\n", userId, err)
	}

	*reply = true

	return nil
}

func (r *RPCServer) UpdatePassord() error {
	return nil
}