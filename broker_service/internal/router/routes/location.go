package routes

import "net/http"

const LOCATION_RESOURCE = "/location"

var LocationRoutes = []*Route{
	{
		Uri:    "",
		Method: "",
		Function: func(http.ResponseWriter, *http.Request) {
		},
		Authentication: false,
	},
}
