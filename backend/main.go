package main

import (
	"cms/backend/route"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func newApp() *iris.Application {
	app := iris.New()
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	// route.RegisterRoute(app)
	// mvc.Configure(app.Party("/users"), route.RegisterRoute)
	route.RegisterRoute(app)

	return app
}

func main() {
	app := newApp()

	// route.RegisterRoute(app)

	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("resources", ".html"))
	app.StaticWeb("/static", "resources/static") // 设置静态资源

	app.Run(iris.Addr(":8080"))
}
