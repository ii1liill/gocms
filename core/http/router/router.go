package router

import (
	"github.com/julienschmidt/httprouter"
	"sync"
)

type GroupStack struct {
	Prefix string
}

type Router struct {
	GroupStack []GroupStack
	HttpRouter *httprouter.Router
}

func (router *Router) UpdateGroupStack(GroupStack GroupStack) {
	router.GroupStack = append(router.GroupStack, GroupStack)
}

// Group 利用匿名函数传递router
func (router Router) Group(GroupStack GroupStack, Callback func(router Router)) {
	router.UpdateGroupStack(GroupStack)
	Callback(router)
}

func (router Router) Prefix(prefix string, Callback func(router Router)) {
	GroupStack := GroupStack{Prefix: prefix}
	router.UpdateGroupStack(GroupStack)
	Callback(router)
}

// Handle 在执行httprouter的Handle之前先处理GroupStack
func (router *Router) Handle(method string, path string, handle httprouter.Handle) {
	// prefix
	var prefix string
	// 处理groupStack
	for _, groupStack := range router.GroupStack {
		prefix += groupStack.Prefix
	}
	path = prefix + path
	router.HttpRouter.Handle(method, path, handle)
}

func (router *Router) DELETE(path string, handle httprouter.Handle) {
	router.Handle("DELETE", path, handle)
}
func (router *Router) GET(path string, handle httprouter.Handle) {
	router.Handle("GET", path, handle)
}
func (router *Router) HEAD(path string, handle httprouter.Handle) {
	router.Handle("HEAD", path, handle)
}
func (router *Router) OPTIONS(path string, handle httprouter.Handle) {
	router.Handle("OPTIONS", path, handle)
}
func (router *Router) PATCH(path string, handle httprouter.Handle) {
	router.Handle("PATCH", path, handle)
}
func (router *Router) POST(path string, handle httprouter.Handle) {
	router.Handle("POST", path, handle)
}
func (router *Router) PUT(path string, handle httprouter.Handle) {
	router.Handle("PUT", path, handle)
}


// 利用sync.Once方法实现单例模式生成Router对象
var router *Router
var once sync.Once

func New() *Router {
	once.Do(func() {
		router = &Router {
			HttpRouter: httprouter.New(),
		}
	})
	return router
}
