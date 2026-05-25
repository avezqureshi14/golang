package gateway

import (
	"net/http/httputil"

	gateway "gateway/internal/proxy"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Path      string
	Methods   []string
	Proxy     *httputil.ReverseProxy
	Protected bool
}

func Register(r *gin.Engine, jwtMiddleware gin.HandlerFunc, routes []Route) {
	for _, route := range routes {

		handler := gateway.Handler(route.Proxy)

		var finalHandler gin.HandlerFunc
		if route.Protected {
			finalHandler = func(c *gin.Context) {
				jwtMiddleware(c)
				if c.IsAborted() {
					return
				}
				handler(c)
			}
		} else {
			finalHandler = handler
		}

		for _, method := range route.Methods {

			// handle ANY
			if method == "ANY" {
				r.Any(route.Path, finalHandler)
				continue
			}

			// normal methods
			r.Handle(method, route.Path, finalHandler)
		}
	}
}