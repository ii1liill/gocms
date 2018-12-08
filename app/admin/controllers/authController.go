package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthController struct {

}

var Auth = AuthController{}

func (c *AuthController) Login(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {

	fmt.Fprintf(writer,"请您登录")
}
