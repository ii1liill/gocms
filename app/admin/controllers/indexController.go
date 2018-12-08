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

func (c *IndexController) Index(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {

	fmt.Fprintf(writer,"hello world")
}