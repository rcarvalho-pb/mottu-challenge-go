package routes

import (
	"net/http"
)

const AUTH_RESOURCE = "/auth"

var AuthRoutes = []*Route{
	{
		Uri:            AUTH_RESOURCE,
		Method:         http.MethodPost,
		Function:       controller.AuthController.Authenticate,
		Authentication: false,
	},
}
