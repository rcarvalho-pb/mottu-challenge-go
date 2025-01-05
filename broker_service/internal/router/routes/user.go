package routes

import "net/http"

const USER_RESOURCE = "/user"

var UserRoutes = []*Route{
	{
		Uri:    "",
		Method: "",
		Function: func(http.ResponseWriter, *http.Request) {
		},
		Authentication: false,
	},
}
