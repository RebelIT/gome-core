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
		log.Printf("ERROR: roku HandlerInfoGet from DB %s", err)
		return
	}

	info, err := c.getInfo()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		log.Printf("ERROR: roku HandlerInfoGet dev info %s", err)
		return
	}

	respond(w, r, xmlToJsonInfo(info))
	return
}

func HandlerPowerSet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := strings.ToLower(vars["name"])
	power := strings.ToLower(vars["state"])

	//todo: validate the input from defined arr of options in powerToBool func.
	state, err := powerToBool(power)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusBadRequest)
		log.Printf("ERROR: roku HandlerPowerSet %s", err)
		return
	}

	c, err := GetDeviceFromDb(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		log.Printf("ERROR: roku HandlerPowerSet %s", err)
		return
	}

	if err := c.controlPowerState(state); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		log.Printf("ERROR: roku HandlerPowerSet %s", err)
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
		log.Printf("ERROR: roku HandlerPowerGet %s", err)
		return
	}

	state, err := c.getPowerState()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		log.Printf("ERROR: roku HandlerPowerGet %s", err)
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
		log.Printf("ERROR: roku HandlerOnlineGet %s", err)
		return
	}

	state, err := c.getOnlineState()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		log.Printf("ERROR: roku HandlerOnlineGet %s", err)
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
		log.Printf("ERROR: roku HandlerAppGet %s", err)
		return
	}

	apps, err := c.getApps()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		log.Printf("ERROR: roku HandlerAppGet %s", err)
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
		log.Printf("ERROR: roku HandlerAppActiveGet %s", err)
		return
	}

	app, err := c.getActiveApp()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		log.Printf("ERROR: roku HandlerAppActiveGet %s", err)
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
		log.Printf("ERROR: roku HandlerApplaunch %s", err)
		return
	}

	apps, err := c.getApps()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		log.Printf("ERROR: roku HandlerApplaunch %s", err)
		return
	}

	if !validateAppInput(id, apps) {
		w.WriteHeader(http.StatusBadRequest)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusBadRequest)
		log.Printf("ERROR: roku HandlerApplaunch %s", err)
		return
	}

	if err := c.launchApp(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		log.Printf("ERROR: roku HandlerApplaunch %s", err)
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
		log.Printf("ERROR: roku HandlerKeypress %s", fmt.Errorf("bad input %s", key))
		return
	}

	c, err := GetDeviceFromDb(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		log.Printf("ERROR: roku HandlerKeypress %s", err)
		return
	}

	if err := c.keyPress(key); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		stat.Http(r.Method, stat.HTTPIN, r.URL.String(), http.StatusInternalServerError)
		log.Printf("ERROR: roku HandlerKeypress %s", err)
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
			Text:    a.Text,
			Version: a.Version,
		}
		jApps.Apps = append(jApps.Apps, ja)
	}

	return jApps
}

func xmlToJsonActiveApp(app ActiveApp) JsonApp {
	return JsonApp{
		ID:      app.App.ID,
		Text:    app.App.Text,
		Version: app.App.Version,
	}
}
