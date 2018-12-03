package gocms

import (
	"github.com/ii1liill/gocms/core"
	_ "github.com/ii1liill/gocms/app/routes"
)
// 配置路径

func Run(configPath string) {
	app := core.New(configPath)
	app.Run()
}
