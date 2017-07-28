package server

import (
	"github.com/jinzhu/configor"
	"gopkg.in/mgo.v2"
	"log"
	"github.com/bxyb214/iot-lock-server/model"
)

var Config = struct {

	APPName          string `default:"iot-lock-server"`

	Server struct {
		Port         string   `default:"8080"`
	}
	Api struct {
		Prefix       string   `default:"/api"`
	}
	DB struct {
		Url          string   `required:"true"`
	}

}{}

func init() {

	// Initialize AppConfig variable
	initConfig()
	// Start a MongoDB session
	initDbSession()
}

func initConfig()  {
	configor.Load(&Config, "config.yml")
}

//create database session
func initDbSession()  {

	session, err := mgo.Dial(Config.DB.Url)
	if err != nil {
		log.Fatalf("[createDbSession]: %s\n", err)
	}
	session.SetMode(mgo.Monotonic, true)

	model.DB = session.DB("iot")
}
