package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func OrdersProxy() http.Handler {
	target, _ := url.Parse("http://service_orders:8002")
	return httputil.NewSingleHostReverseProxy(target)
}
