package routes

import (
	"fmt"
	"net/http"
)

type Route struct {
	Uri            string
	Method         string
	Function       func(http.ResponseWriter, *http.Request)
	Authentication bool
}

func config(mux *http.ServeMux) *http.ServeMux {
	var routes []*Route
	routes = append(routes, AuthRoutes...)
	routes = append(routes, UserRoutes...)
	routes = append(routes, LocationRoutes...)
	routes = append(routes, MotorcycleRoutes...)
	for _, route := range routes {
		if route.Authentication {
			mux.HandleFunc(fmt.Sprintf("%s %s", route.Method, route.Uri), route.Function)
		} else {

		}
	}

	return mux
}
