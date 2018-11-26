package ocms

import (
	"fmt"

	myhttp "github.com/ii1liill/gocms/http"
)

// 常用常量
const (
	defaultAppPath = "github.com/ii1liill/gocms/app"
)

// OCms 定义核心结构体
type OCms struct {
	// 应用配置项
	Config interface{}
	// 应用默认目录
	AppDir string
	// web服务运行绑定地址
	BindAddr string
}

// New 生成OCms对象的函数
func New() OCms {
	// fmt.Println(config.Get("tests.testsInt"))
	o := OCms{
		AppDir:   defaultAppPath,
		BindAddr: ":8000",
	}

	return o
}

// Run OCms的启动方法
func (o *OCms) Run() error {

	// 启动http服务器
	fmt.Println("Service running!")
	return myhttp.Bind(o.BindAddr, nil)
}
