package gateway

import "net/http/httputil"

func UserRoutes(userService *httputil.ReverseProxy) []Route {
	return []Route{
		{
			Path:      "/users/*path",
			Methods:   []string{"ANY"},
			Proxy:     userService,
			Protected: true,
		},
	}
}