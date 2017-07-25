package main

import (
	"github.com/jinzhu/configor"
	"github.com/kataras/iris"
)

var Config = struct {

	APPName          string `default:"iot-lock-server"`

	Server struct {
		Port         string   `default:"8080"`
	}
	Api struct {
		Prefix       string   `default:"/api"`
	}

}{}


func main() {
	configor.Load(&Config, "config.yml")

	app := iris.New()

	//TODO how to enable iris debug

	Route(app)

	app.Run(iris.Addr(":" + Config.Server.Port), iris.WithCharset("UTF-8"))

}

