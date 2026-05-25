package gateway

import "net/http/httputil"

func TodoRoutes(todoService *httputil.ReverseProxy) []Route {
	return []Route{
		{
			Path:      "/todos/*path",
			Methods:   []string{"GET"},
			Proxy:     todoService,
			Protected: false,
		},
		{
			Path:      "/todos/*path",
			Methods:   []string{"POST", "PUT", "DELETE", "PATCH"},
			Proxy:     todoService,
			Protected: true,
		},
	}
}