package gateway

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewProxy(target string) *httputil.ReverseProxy {
	url, _ := url.Parse(target)

	proxy := httputil.NewSingleHostReverseProxy(url)

	originalDirector := proxy.Director

	proxy.Director = func(req *http.Request) {
		originalDirector(req)
	}

	return proxy
}