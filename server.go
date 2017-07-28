package server

import (
	"github.com/kataras/iris"
)

func StartServer() {

	app := iris.New()

	//TODO how to enable iris debug
	Route(app)

	app.Run(iris.Addr(":" + Config.Server.Port), iris.WithCharset("UTF-8"))
}
