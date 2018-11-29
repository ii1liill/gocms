package gocms

import (
	"fmt"
	"github.com/ii1liill/gocms/config"

	myhttp "github.com/ii1liill/gocms/http"
)

// 常用常量
const (
	defaultAppPath = "github.com/ii1liill/gocms/app"
)

// GoCms 定义核心结构体
type GoCms struct {
	// 应用配置项
	Config interface{}
	// 应用默认目录
	AppDir string
	// web服务运行绑定地址
	BindAddr string
}

// New 生成OCms对象的函数
func New(configPath string) GoCms {
	config.SetConfigPath(configPath)
	o := GoCms{
		AppDir:   defaultAppPath,
		BindAddr: ":8000",
	}

	return o
}

// Run GoCms的启动方法
func (o *GoCms) Run() error {

	// 启动http服务器
	fmt.Println("Service running!")
	return myhttp.Bind(o.BindAddr, nil)
}
