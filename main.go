package main

import (
	"github.com/kataras/iris"
	"hk_shopping/conf"
	"hk_shopping/filter"
	"hk_shopping/models"
)

func main() {
	app := iris.New()
	config(app)
	app.Run(iris.Addr(conf.Config.Addr), iris.WithConfiguration(conf.ConfigTOML))
}

//相关配置
func config(app *iris.Application) {
	app.Use(filter.Before)
	app.Done(filter.After)

	//服务关闭时清除资源
	app.ConfigureHost(func(h *iris.Supervisor) {
		h.RegisterOnShutdown(func() {
			models.Close()
		})
	})
}
