package gateway

import "net/http/httputil"

func AuthRoutes(userService *httputil.ReverseProxy) []Route {
	return []Route{
		{
			Path:      "/auth/*path",
			Methods:   []string{"ANY"},
			Proxy:     userService,
			Protected: false,
		},
	}
}