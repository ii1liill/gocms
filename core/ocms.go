package ocms

import (
	"fmt"

	ochttp "github.com/ii1liill/gocms/core/http"
)

const (
	defaultAppPath = "../app"
)

type OCms struct {
	// 应用配置项
	Config interface{}
	// 应用默认目录
	AppDir string
	// web服务运行绑定地址
	BindAddr string
}

func New() OCms {
	o := OCms{
		AppDir:   defaultAppPath,
		BindAddr: ":8000",
	}

	return o
}

func (o *OCms) Run() error {

	// 启动http服务器
	fmt.Println("Service running!")
	return ochttp.Bind(o.BindAddr, nil)
}
