package main

import (
	"github.com/rebelit/gome-core/common/config"
	"github.com/rebelit/gome-core/core/devices"
	"github.com/rebelit/gome-core/core/web"
	"log"
	"net/http"
)

func main() {
	config.Runtime()
	devices.InitializeDatabases()

	start(config.App.ListenPort)

	return
}

func start(port string) {
	log.Printf("Starting %s on http:%s", config.App.Name, port)
	router := web.NewRouter()
	log.Fatal(http.ListenAndServe(":"+port, router))
}
