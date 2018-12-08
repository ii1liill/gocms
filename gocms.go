package gocms

import (
	"github.com/ii1liill/gocms/core"
)
// 配置路径

func Run(configPath string) {
	app := core.New(configPath)
	app.Run()
}
