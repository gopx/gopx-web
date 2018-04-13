package main

import (
	"runtime"

	"gopx.io/gopx-web/pkg/log"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// log.SetFormatter(&log.TextFormatter{})

	// log.WithFields(log.Fields{
	// 	"animal": "walrus",
	// }).Info("A walrus appears")

	// log.WithFields(log.Fields{
	// 	"animal": "walrus",
	// }).Info("The server is running")

	// log.WithFields(log.Fields{
	// 	"animal": "walrus",
	// }).Warn("A walrus appears")

	// log.WithFields(log.Fields{
	// 	"animal": "walrus",
	// }).

	// log.Warn("sdfs")

	// log.Panic("SDf")

	// log.Fatal(http.ListenAndServe(":8080", route.NewGoPXRouter()))

	log.Debug("Server", "Useful debugging information.")
	log.Info("IndexRouting", "Something noteworthy happened!")
	log.Warn("Templating", "You should probably take a look at this.")
	log.Error("Unknwon", "Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	//log.Fatal("Bye.")
	// Calls panic() after logging
	//log.Panic("About page", "I'm bailing.")

}
