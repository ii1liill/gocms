package http

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type Handle httprouter.Handle

func Bind(addr string, handler http.Handler) error {
	return http.ListenAndServe(addr, handler)
}
