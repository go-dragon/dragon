package dragon

import (
	"dragon/core/dragon/conf"
	"dragon/core/dragon/dlogger"
	"dragon/core/dragon/dnacos"
	"dragon/core/dragon/dragonzipkin"
	"dragon/core/dragon/dredis"
	"dragon/domain/repository"
	"dragon/tools/dmongo"
	"github.com/spf13/viper"
	"log"
	"net/http"
	_ "net/http/pprof"
)

// AppInit func
// do some components init
// todo add Prometheus in the future
func AppInit() {
	//init config
	conf.InitConf()

	// init pprof
	if viper.GetBool("server.pprof.enable") {
		var host string
		if viper.GetString("server.pprof.host") != "" {
			host = viper.GetString("server.pprof.host")
		} else {
			host = conf.IntranetIp
		}
		go func() {
			log.Println("Pprof server on "+host+":"+viper.GetString("server.pprof.port"), "http://"+host+":"+viper.GetString("server.pprof.port")+"/debug/pprof")
			http.ListenAndServe(host+":"+viper.GetString("server.pprof.port"), nil)
		}()
	}

	// init zipkin server middleware
	if viper.GetBool("zipkin.enable") {
		dragonzipkin.Init()
	}

	// init nacos
	if viper.GetBool("nacos.enable") {
		dnacos.Init()
	}

	//init db
	repository.InitDB()

	//init redis
	dredis.InitRedis()

	// init mongodb
	dmongo.InitDB()

	// init logger
	go func() {
		dlogger.TickFlush()
	}()
}
