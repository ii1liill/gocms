package middleware

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func AdminLoginAuth(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		next(writer, request, params)
	}
}