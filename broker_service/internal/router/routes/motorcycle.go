package routes

import "net/http"

const MOTORCYCLE_RESOURCE = "/motorcycle"

var MotorcycleRoutes = []*Route{
	{
		Uri:    "",
		Method: "",
		Function: func(http.ResponseWriter, *http.Request) {
		},
		Authentication: false,
	},
}
