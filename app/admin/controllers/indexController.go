package controllers

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	)

type IndexController struct {

}
// 供routes调用
var Index = IndexController{}

func (c *IndexController) Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w,"hello world")
}