package main

import (
	"dragon/core/dragon"
	"dragon/middleware"
	"dragon/router"
)

func main() {
	//dragon init conf, pprof mysql, redis mongodb and so on
	dragon.AppInit()
	//init dragon
	dr := dragon.New()
	//init route, you can chain any middleware here :)
	//dr.InitRoute(dragonzipkin.ServerMiddleware(middleware.LogInfo(router.Routes)))
	dr.InitRoute(middleware.LogInfo(router.Routes))
	//dragon fly
	dr.Fly()
}
