package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/ii1liill/gocms/core/session"
)

func AdminLoginAuth(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		sess,_ := session.Store.Get(request, "cms_admin")
		_, exists := sess.Values["user_id"]
		if !exists {
			writer.Header().Set("Location", "/admin/login")
			writer.WriteHeader(301)
			return
		}
		next(writer, request, params)
	}
}