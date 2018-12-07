package routes

import (
	"github.com/ii1liill/gocms/app/admin/controllers"
	"github.com/ii1liill/gocms/app/http/middleware"
	"github.com/ii1liill/gocms/core/http/router"
)

func init() {
	r := router.New()

	r.Prefix("/admin", func(r router.Router) {
		r.Middleware(middleware.AdminLoginAuth, func(r router.Router) {
			r.GET("/", controllers.Index.Index)
			r.Prefix("/test", func(r router.Router) {
				r.GET("/a", controllers.Index.Index)
			})
			r.GET("/hello", controllers.Index.Index)
		})
	})
}