package routes

import (
	"fmt"

	"github.com/ii1liill/gocms/core/config"
	"github.com/ii1liill/gocms/core/http/router"
	"github.com/ii1liill/gocms/core/session"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Web() {
	r := router.New()
	r.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		sess,_ := session.Store.Get(request, "cms_admin")
		fmt.Println(sess.Values["name"])
		fmt.Fprintf(writer, "hello %s", config.Get("app.name"))
	})
}