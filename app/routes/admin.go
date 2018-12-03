package routes

import (
	"fmt"
	"github.com/ii1liill/gocms/app/admin/controllers"
	"github.com/ii1liill/gocms/core/http/router"
)

func init() {
	r := router.New()
	adminGroupStack := router.GroupStack{
		"prefix": "/admin",
	}
	r.Group(adminGroupStack, func(r router.Router) {
		fmt.Println(r)
		r.GET("/", controllers.Index.Index)
	})
}