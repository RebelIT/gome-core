package roku

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rebelit/gome-core/common/stat"
	"log"
	"net/http"
	"strings"
)

//HandlerInfoGet
func HandlerInfoGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := strings.ToLower(vars["name"])

	c, err := GetDeviceFromDb(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		return
	}

	info, err := c.getInfo()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		return
	}

	respond(w, r, xmlToJsonInfo(info))
	return
}

func HandlerPowerSet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := strings.ToLower(vars["name"])
	power := strings.ToLower(vars["state"])

	state, err := powerToBool(power)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusBadRequest)
		return
	}

	c, err := GetDeviceFromDb(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		return
	}

	if err := c.controlPowerState(state); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusOK)
	return
}

func HandlerPowerGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := strings.ToLower(vars["name"])

	c, err := GetDeviceFromDb(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		return
	}

	state, err := c.getPowerState()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		return
	}

	response := RespPwr{
		Name:  c.Name,
		State: state,
	}

	respond(w, r, response)
	return
}

func HandlerOnlineGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := strings.ToLower(vars["name"])

	c, err := GetDeviceFromDb(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		return
	}

	state, err := c.getOnlineState()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		return
	}

	response := RespPwr{
		Name:  c.Name,
		State: state,
	}

	respond(w, r, response)
	return
}

func HandlerAppGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := strings.ToLower(vars["name"])

	c, err := GetDeviceFromDb(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		return
	}

	apps, err := c.getApps()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		return
	}

	respond(w, r, xmlToJsonApps(apps))
	return
}

func HandlerAppActiveGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := strings.ToLower(vars["name"])

	c, err := GetDeviceFromDb(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		return
	}

	app, err := c.getActiveApp()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		return
	}

	respond(w, r, xmlToJsonActiveApp(app))
	return
}

func HandlerApplaunch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := strings.ToLower(vars["name"])
	id := vars["id"]

	c, err := GetDeviceFromDb(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		return
	}

	apps, err := c.getApps()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		return
	}

	if !validateAppInput(id, apps) {
		w.WriteHeader(http.StatusBadRequest)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusBadRequest)
		return
	}

	if err := c.launchApp(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func HandlerKeypress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := strings.ToLower(vars["name"])
	key := strings.ToLower(vars["key"])

	if !validateKeyInput(key) {
		w.WriteHeader(http.StatusBadRequest)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusBadRequest)
		return
	}

	c, err := GetDeviceFromDb(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		return
	}

	if err := c.keyPress(key); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

//private handler helpers
func powerToBool(state string) (boolState bool, error error) {
	if state == "on" {
		return true, nil
	}
	if state == "off" {
		return false, nil
	}

	return false, fmt.Errorf("bad input")
}

func validateAppInput(id string, apps Apps) bool {
	for _, app := range apps.App {
		if app.ID == id {
			return true
		}
	}

	return false
}

func validateKeyInput(key string) bool {
	for _, k := range keys {
		if key == k {
			return true
		}
	}

	return false
}

func respond(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Printf("ERROR: roku handler response, %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
			return
		}
	}

	stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusOK)
	return
}

func xmlToJsonInfo(info DeviceInfo) JsonDeviceInfo {
	return JsonDeviceInfo{
		Udn:                info.Udn,
		SerialNumber:       info.SerialNumber,
		DeviceID:           info.DeviceID,
		IsTv:               info.IsTv,
		IsStick:            info.IsStick,
		NetworkName:        info.NetworkName,
		UserDeviceName:     info.UserDeviceName,
		UserDeviceLocation: info.UserDeviceLocation,
		Uptime:             info.Uptime,
		PowerMode:          info.PowerMode,
	}

}

func xmlToJsonApps(apps Apps) JsonApps {
	jApps := JsonApps{}

	for _, a := range apps.App {
		ja := JsonApp{
			ID:      a.ID,
			Type:    a.Type,
			Version: a.Version,
		}
		jApps.Apps = append(jApps.Apps, ja)
	}

	return jApps
}

func xmlToJsonActiveApp(app App) JsonApp {
	return JsonApp{
		ID:      app.ID,
		Type:    app.Type,
		Version: app.Version,
	}
}
