package dragonzipkin

import (
	"errors"
	zipkin "github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
	"github.com/openzipkin/zipkin-go/reporter"
	reporterhttp "github.com/openzipkin/zipkin-go/reporter/http"
	"github.com/spf13/viper"
	"log"
	"net"
	"net/http"
)

var Tracer *zipkin.Tracer
var ServerMiddleware func(handler http.Handler) http.Handler
var Reporter reporter.Reporter
var Client *zipkinhttp.Client

// InitServerMiddleware and client
func Init() {
	// set up a span reporter
	Reporter = reporterhttp.NewReporter(viper.GetString("zipkin.reporter"))
	//defer reporter.Close()
	// create our local service endpoint
	ip, err := getClientIp()
	if err != nil {
		log.Fatalf("unable to get ClientIp: %+v\n", err)
	}
	endpoint, err := zipkin.NewEndpoint(viper.GetString("server.servicename"), ip+":"+viper.GetString("server.port"))
	if err != nil {
		log.Fatalf("unable to create local endpoint: %+v\n", err)
	}
	// initialize our tracer
	Tracer, err = zipkin.NewTracer(Reporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		log.Fatalf("unable to create tracer: %+v\n", err)
	}
	// create global zipkin http server middleware
	ServerMiddleware = zipkinhttp.NewServerMiddleware(
		Tracer, zipkinhttp.TagResponseSize(true),
	)

	// create global zipkin traced http client
	Client, err = zipkinhttp.NewClient(Tracer, zipkinhttp.ClientTrace(true))
	if err != nil {
		log.Fatalf("unable to create client: %+v\n", err)
	}
}

func getClientIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", errors.New("Can not find the client ip address!")

}
