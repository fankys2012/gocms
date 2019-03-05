package route

import (
	"cms/backend/controllers"

	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
)

//注册路由，如果失败返回错误
// func RegisterRoute(app *mvc.Application) {

// 	app.Handle(new(controllers.UserController))
// }

func RegisterRoute(app *iris.Application) {

	app.PartyFunc("/user", func(r iris.Party) {
		r.Get("/login", hero.Handler(controllers.NewUserController().Login))
		r.Get("/json", hero.Handler(controllers.NewUserController().Json))
	})

	app.PartyFunc("/inject", func(inject iris.Party) {
		controller := controllers.NewInjectController()
		inject.Post("/assetsInject", hero.Handler(controller.AssetsInject))
	})
}
