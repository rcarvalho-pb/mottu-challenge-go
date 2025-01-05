package routes

import (
	"net/http"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/controllers"
)

const AUTH_RESOURCE = "/auth"

var AuthRoutes = []*Route{
	{
		Uri:            AUTH_RESOURCE,
		Method:         http.MethodPost,
		Function:       controllers.Authenticate,
		Authentication: false,
	},
}
