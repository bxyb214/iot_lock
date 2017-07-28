package server

import (
    "github.com/kataras/iris"
	"github.com/bxyb214/iot-lock-server/apis"
)

func Route(app *iris.Application) {

	apiPrefix   := Config.Api.Prefix

	router := app.Party(apiPrefix)
	{
		router.Get("/login", apis.Login)
		router.Post("/carports", apis.CreateCarport)
		router.Get("/carports", apis.GetAllCarports)
	}
}
