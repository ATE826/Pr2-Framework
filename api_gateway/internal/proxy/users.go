package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func UsersProxy() http.Handler {
	target, _ := url.Parse("http://service_users:8001")
	return httputil.NewSingleHostReverseProxy(target)
}
