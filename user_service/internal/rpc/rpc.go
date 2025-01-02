package rpc

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/rcarvalho-pb/mottu-user_service/internal/application/dtos"
	"github.com/rcarvalho-pb/mottu-user_service/internal/application/services"
)

type RPCServer struct {
	userService *services.UserService
	Port        string
}

func New(service *services.UserService, port string) *RPCServer {
	return &RPCServer{
		userService: service,
		Port:        port,
	}
}

func (r *RPCServer) RPCListen() error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", r.Port))
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

func (r *RPCServer) CreateUser(newUser *dtos.UserDTO, _ *struct{}) error {
	if newUser == nil {
		return fmt.Errorf("user can't be null")
	}

	if err := r.userService.CreateUser(newUser); err != nil {
		return err
	}

	return nil
}

func (r *RPCServer) GetUserById(userId *int64, reply *dtos.UserDTO) error {
	user, err := r.userService.GetUserById(*userId)
	if err != nil {
		return fmt.Errorf("error getting user by id: %s\n", err)
	}

	*reply = *user

	return err
}

func (r *RPCServer) GetUserByUsername(username *string, reply *dtos.UserDTO) error {
	user, err := r.userService.GetUserByUsername(*username)
	if err != nil {
		return fmt.Errorf("error getting user by username: %s\n", err)
	}

	*reply = *user
	return err
}

func (r *RPCServer) GetAllActiveUsers(_ struct{}, reply *[]*dtos.UserDTO) error {
	users, err := r.userService.GetAllActiveUsers()
	if err != nil {
		return fmt.Errorf("error getting all active users: %s\n", err)
	}

	*reply = users
	return err
}

func (r *RPCServer) GetAllUsers(_ struct{}, reply *[]*dtos.UserDTO) error {
	users, err := r.userService.GetAllUsers()
	if err != nil {
		return fmt.Errorf("error getting all users: %s\n", err)
	}

	*reply = users
	return err
}

func (r *RPCServer) DeactivateUser(userId *int64, reply *bool) error {
	if err := r.userService.DeactivateUserById(*userId); err != nil {
		return fmt.Errorf("error deactivating user [%d]: %s\n", userId, err)
	}

	*reply = true

	return nil
}

func (r *RPCServer) ReactivateUser(userId int64, _ *struct{}) error {
	if err := r.userService.ReactivateUserById(userId); err != nil {
		return fmt.Errorf("error reactivating user [%d]: %s\n", userId, err)
	}

	return nil
}

func (r *RPCServer) ComparePasswords(passwords *dtos.ComparePasswordsDTO, _ *struct{}) error {
	if err := r.userService.ComparePasswords(passwords.HashedPassword, passwords.Password); err != nil {
		return fmt.Errorf("Passwords doesn't match")
	}

	return nil
}

func (r *RPCServer) UpdatePassord(user *dtos.UserDTO, _ *struct{}) error {
	if err := r.userService.UpdateUser(user); err != nil {
		return err
	}

	return nil
}
