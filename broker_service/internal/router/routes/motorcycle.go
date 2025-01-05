package routes

import (
	"fmt"
	"net/http"
)

const MOTORCYCLE_RESOURCE = "/motorcycle"

var MotorcycleRoutes = []*Route{
	{
		Uri:            fmt.Sprintf("%s/{id}", MOTORCYCLE_RESOURCE),
		Method:         http.MethodGet,
		Function:       controller.MotorcycleController.GetMotorcycleById,
		Authentication: true,
		Admin:          false,
	},
	{
		Uri:            fmt.Sprintf("%s/get", MOTORCYCLE_RESOURCE),
		Method:         http.MethodPost,
		Function:       controller.MotorcycleController.GetAllMotorcycles,
		Authentication: true,
		Admin:          false,
	},
	{
		Uri:            fmt.Sprintf("%s/get/{year}", MOTORCYCLE_RESOURCE),
		Method:         http.MethodGet,
		Function:       controller.MotorcycleController.GetMotorcyclesByYear,
		Authentication: true,
		Admin:          false,
	},
}
