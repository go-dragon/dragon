package dragon

import (
	"dragon/core/dragon/conf"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type Dragon struct {
	Handler http.Handler
}

//new dragon
func New() *Dragon {
	return new(Dragon)
}

// init route
func (dragon *Dragon) InitRoute(handler http.Handler) {
	dragon.Handler = handler
}

//start listening
func (dragon *Dragon) Fly() {

	//dragon fly
	log.Println("env: " + conf.Env)
	log.Println("set environment variable DRAGON dev,test,prod 🐲🐲🐲")
	if conf.IntranetIp == "" {
		conf.IntranetIp = "127.0.0.1"
	}
	webAddr := "http://" + conf.IntranetIp + ":" + viper.GetString("server.port")
	if viper.GetString("server.host") != "" {
		webAddr = "http://" + viper.GetString("server.host") + ":" + viper.GetString("server.port")
	}
	log.Println("start server on: " + webAddr)
	log.Fatal(http.ListenAndServe(viper.GetString("server.host")+":"+viper.GetString("server.port"), dragon.Handler))
}
