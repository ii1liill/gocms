package router

import (
	"net/http"
)

func Get(r string, handle func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(r, handle)
}

func Post(r string, handle interface{}) {

}

func Any() {

}

func Bind(o interface{}) {

}
