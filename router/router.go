package router

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"hk_shopping/controller"
)

//路由
func Route(app *iris.Application) {
	mvc.Configure(app.Party("/user"), func(m *mvc.Application) {
		m.Handle(&controller.UserController{})
	})

	mvc.Configure(app.Party("/ad"), func(m *mvc.Application) {
		m.Handle(&controller.AdController{})
	})
}
