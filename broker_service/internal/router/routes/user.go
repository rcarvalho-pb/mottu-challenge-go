package routes

import (
	"fmt"
	"net/http"
)

const USER_RESOURCE = "/user"

var UserRoutes = []*Route{
	{
		Uri:            fmt.Sprintf("%s/id/{id}", USER_RESOURCE),
		Method:         http.MethodGet,
		Function:       controller.UserController.GetUserById,
		Authentication: true,
	},
	{
		Uri:            fmt.Sprintf("%s/username/{id}", USER_RESOURCE),
		Method:         http.MethodGet,
		Function:       controller.UserController.GetUserByUsername,
		Authentication: true,
	},
	{
		Uri:            fmt.Sprintf("%s/update", USER_RESOURCE),
		Method:         http.MethodPost,
		Function:       controller.UserController.UpdateUser,
		Authentication: true,
	},
	{
		Uri:            fmt.Sprintf("%s", USER_RESOURCE),
		Method:         http.MethodGet,
		Function:       controller.UserController.GetAllUsers,
		Authentication: true,
	},
	{
		Uri:            fmt.Sprintf("%s/new", USER_RESOURCE),
		Method:         http.MethodPost,
		Function:       controller.UserController.CreateUser,
		Authentication: false,
	},
}
