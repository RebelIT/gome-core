package web

import (
	"github.com/gorilla/mux"
	"github.com/rebelit/gome-core/common/config"
	"github.com/rebelit/gome-core/common/stat"
	"github.com/rebelit/gome-core/core/devices"
	"github.com/rebelit/gome-core/core/devices/roku"
	"log"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
		router.Use(authMiddleware)
	}

	return router
}

var routes = Routes{
	Route{"status", "GET", "/api/status", status},
	Route{"device", "GET", "/api/device/{type}", devices.GetDevices},
	Route{"device", "POST", "/api/device/{type}", devices.LoadDevice},
	Route{"device", "GET", "/api/deviceTypes", devices.GetDeviceTypes},
	Route{"roku", "GET", "/api/roku/{name}/info", roku.HandlerInfoGet},
	Route{"roku", "GET", "/api/roku/{name}/online", roku.HandlerOnlineGet},
	Route{"roku", "GET", "/api/roku/{name}/power", roku.HandlerPowerGet},
	Route{"roku", "PUT", "/api/roku/{name}/power/{state}", roku.HandlerPowerSet},
	Route{"roku", "GET", "/api/roku/{name}/app", roku.HandlerAppGet},
	Route{"roku", "GET", "/api/roku/{name}/app/active", roku.HandlerAppActiveGet},
	Route{"roku", "POST", "/api/roku/{name}/app/launch/{id}", roku.HandlerApplaunch},
	Route{"roku", "POST", "/api/roku/{name}/key/{key}", roku.HandlerKeypress},
	//Route{"rpiot", "GET", "/api/rpiot/{name}/info", rpiot.HandlerInfoGet},
}

func status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	stat.Http(r.Method, "inbound", r.URL.String(), http.StatusOK)

	return
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/api/status" {
			//skip auth for health check
			next.ServeHTTP(w, r)

		} else {
			authorization := r.Header.Get("Authorization")
			if validateAuth(authorization) {
				// Pass down the request to the next handler
				log.Printf("INFO: http authorized %s:%s", r.Method, r.URL.String())
				next.ServeHTTP(w, r)

			} else {
				log.Printf("INFO: http unauthorized %s:%s", r.Method, r.URL.String())
				w.WriteHeader(http.StatusUnauthorized)
				stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusUnauthorized)

				return
			}
		}
	})
}

func validateAuth(authorization string) bool {
	if authorization == "Bearer "+config.App.AuthToken {
		return true
	}

	return false
}
