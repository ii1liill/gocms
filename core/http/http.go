package http

import (
	"net/http"
)

func Bind(addr string, handler http.Handler) error {
	return http.ListenAndServe(addr, handler)
}
