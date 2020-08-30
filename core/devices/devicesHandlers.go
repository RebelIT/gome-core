package devices

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rebelit/gome-core/common/stat"
	"github.com/rebelit/gome-core/core/devices/roku"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetDevices(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	device := strings.ToLower(vars["type"])

	if device == "all" {
		devices, err := getAllLoadedDevices(device)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
			return
		}

		respond(w, r, devices)
	}

	devices, err := getAllLoadedDevices(device)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		return
	}

	respond(w, r, devices)
	return
}

func GetDeviceTypes(w http.ResponseWriter, r *http.Request) {
	deviceType := []DeviceType{}

	types, err := getLoadedDeviceTypes()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		return
	}

	for _, t := range types {
		dt := DeviceType{Name: t}
		deviceType = append(deviceType, dt)
	}

	respond(w, r, deviceType)
	return
}

func LoadDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deviceType := strings.ToLower(vars["type"])

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusBadRequest)
		return
	}

	switch strings.ToLower(deviceType) {
	case "roku":
		if err := roku.LoadDevice(data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
			return
		}

	case "rpiot":
		w.WriteHeader(http.StatusNotImplemented)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusNotImplemented)
		return

	default:
		w.WriteHeader(http.StatusBadRequest)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusBadRequest)
		return
	}

	respond(w, r, nil)
	return
}

//helper response
func respond(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Printf("ERROR: devices handler response, %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
			return
		}
	}

	stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusOK)
	return
}
