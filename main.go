package main

import (
	"github.com/rebelit/gome-core/common/config"
	"github.com/rebelit/gome-core/core/devices"
	"github.com/rebelit/gome-core/core/web"
	"log"
	"net/http"
)

func main() {
	log.Printf("INFO: I'm starting")
	config.Runtime()
	devices.InitializeDatabases()

	if config.App.GenerateSpec{
		web.GenerateSpec()
		return
	}

	start(config.App.ListenPort)
	return
}

func start(port string) {
	log.Printf("INFO: listening %s on http:%s", config.App.Name, port)
	router := web.NewRouter()
	log.Fatal(http.ListenAndServe(":"+port, router))
}
