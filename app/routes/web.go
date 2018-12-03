package routes

import (
	"fmt"
	"github.com/ii1liill/gocms/core/config"
	"github.com/ii1liill/gocms/core/http/router"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func init() {
	r := router.New()
	r.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprintf(writer, "hello %s", config.Get("app.name"))
	})
}