package main

import (
	golog "log"
	"net"
	"net/http"
	"os"
	"runtime"
	"strconv"

	"gopx.io/gopx-web/pkg/config"
	"gopx.io/gopx-web/pkg/log"
	"gopx.io/gopx-web/pkg/route"
)

var serverLogger = golog.New(os.Stdout, "", golog.Ldate|golog.Ltime|golog.Lshortfile)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	startServer()
}

func startServer() {
	switch {
	case config.Service.UseHTTP && config.Service.UseHTTPS:
		go startHTTP()
		startHTTPS()
	case config.Service.UseHTTP:
		startHTTP()
	case config.Service.UseHTTPS:
		startHTTPS()
	default:
		log.Fatal("Error: no listener is specified in service config file")
	}
}

func startHTTP() {
	addr := httpAddr()
	router := route.NewGoPXWebRouter()
	server := &http.Server{Addr: addr, Handler: router, ErrorLog: serverLogger}

	log.Info("Running HTTP server on: %s", addr)
	err := server.ListenAndServe()
	log.Fatal("Error: %s", err) // err is always non-nill
}

func startHTTPS() {
	addr := httpsAddr()
	router := route.NewGoPXWebRouter()
	server := &http.Server{Addr: addr, Handler: router, ErrorLog: serverLogger}

	log.Info("Running HTTPS server on: %s", addr)
	err := server.ListenAndServeTLS(config.Service.CertFile, config.Service.KeyFile)
	log.Fatal("Error: %s", err) // err is always non-nill
}

func httpAddr() string {
	return net.JoinHostPort(config.Service.Host, strconv.Itoa(config.Service.HTTPPort))
}

func httpsAddr() string {
	return net.JoinHostPort(config.Service.Host, strconv.Itoa(config.Service.HTTPSPort))
}
