package rpc

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/rcarvalho-pb/mottu-user_service/internal/application/dtos"
	"github.com/rcarvalho-pb/mottu-user_service/internal/application/services"
)

type RPCServer struct {
	userService services.UserService
}

func New(service services.UserService) *RPCServer {
	return &RPCServer{
		userService: service,
	}
}

func (r *RPCServer) RPCListen() error {
	listen, err := net.Listen("tcp", ":12345")
	if err != nil {
		return err
	}

	defer listen.Close()

	err = rpc.RegisterName("UserService", r)
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
	fmt.Println("Here")
	fmt.Println(userId)
	user, err := r.userService.GetUserById(userId)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", user)

	*reply = *user

	return err
}

func (r *RPCServer) GetUsersByUsername(username string, reply *[]*dtos.UserDTO) error {
	users, err := r.userService.GetUsersByUsername(username)
	if err != nil {
		return err
	}

	*reply = users
	return err
}

func (r *RPCServer) GetAllActiveUsers(_ struct{}, reply *[]*dtos.UserDTO) error {
	users, err := r.userService.GetAllActiveUsers()
	if err != nil {
		return err
	}

	*reply = users
	return err
}

func (r *RPCServer) GetAllUsers(_ struct{}, reply *[]*dtos.UserDTO) error {
	users, err := r.userService.GetAllUsers()
	if err != nil {
		return err
	}

	*reply = users
	return err
}
