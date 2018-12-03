package router

import (
	"github.com/julienschmidt/httprouter"
)

type GroupStack map[string]interface{}

type Router struct {
	GroupStack GroupStack
}

var R = httprouter.New()

func (router *Router) UpdateGroupStack(GroupStack GroupStack) {
	router.GroupStack = GroupStack
}

// Group 是局部作用域方法，仅在回调内部起作用
func (router Router) Group(GroupStack GroupStack, Callback func(router Router)) {
	router.UpdateGroupStack(GroupStack)
	Callback(router)
}

// Handle 在执行httprouter的Handle之前先处理GroupStack
func (router *Router) Handle(method string, path string, handle httprouter.Handle) {
	// 处理prefix
	if _, ok := router.GroupStack["prefix"]; ok == true {
		path = router.GroupStack["prefix"].(string) + path
	}
	R.Handle(method, path, handle)
}

func (router *Router) GET(path string, handle httprouter.Handle) {
	router.Handle("GET", path, handle)
}

func New() *Router {
	router := &Router{}
	return router
}

func DELETE(path string, handle httprouter.Handle) {
	R.DELETE(path, handle)
}

func GET(path string, handle httprouter.Handle) {
	R.GET(path, handle)
}

func HEAD(path string, handle httprouter.Handle) {
	R.HEAD(path, handle)
}

func OPTIONS(path string, handle httprouter.Handle) {
	R.OPTIONS(path, handle)
}

func PATCH(path string, handle httprouter.Handle) {
	R.PATCH(path, handle)
}

func POST(path string, handle httprouter.Handle) {
	R.POST(path, handle)
}

func PUT(path string, handle httprouter.Handle) {
	R.PUT(path, handle)
}
