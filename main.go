package main

import (
	"github.com/rebelit/gome-core/web"
	"log"
	"net/http"
)

func main() {
	port := "6660"
	start(port)

	return
}

func start(port string) {
	log.Printf("Starting Web Listener on :%v\n", port)
	router := web.NewRouter()
	log.Fatal(http.ListenAndServe(":"+port, router))
}
